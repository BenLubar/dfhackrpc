package dfhackrpc

import (
	"reflect"

	"github.com/BenLubar/dfhackrpc/dfproto/core"
	"github.com/golang/protobuf/proto"
)

const bindMethod = "BindMethod"
const runCommand = "RunCommand"

type methodInfo struct {
	request  reflect.Type
	response reflect.Type
	id       int16
}

func (c *Client) resetMethods() {
	c.methods = map[[2]string]*methodInfo{
		{"", bindMethod}: {
			request:  reflect.TypeOf(&dfproto_core.CoreBindRequest{}),
			response: reflect.TypeOf(&dfproto_core.CoreBindReply{}),
			id:       0,
		},
		{"", runCommand}: {
			request:  reflect.TypeOf(&dfproto_core.CoreRunCommandRequest{}),
			response: reflect.TypeOf(&dfproto_core.EmptyMessage{}),
			id:       1,
		},
	}
}

// Call runs an RPC method. It is an error to call a method with the wrong
// request or response type. If request or response are nil or their types
// do not match the types used the first time the method was called, Call
// will panic.
func (c *Client) Call(method string, request, response proto.Message) (CommandResult, error) {
	return c.CallPlugin("", method, request, response)
}

// CallPlugin runs an RPC method, similar to Call, but also allows you to
// specify a plugin name in order to call plugin methods API methods.
func (c *Client) CallPlugin(plugin, method string, request, response proto.Message) (CommandResult, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if err := c.checkReconnect(); err != nil {
		return LinkFailure, err
	}

	return c.call(plugin, method, request, response)
}

func (c *Client) bind(plugin, method string, request, response proto.Message) (int16, CommandResult, error) {
	if request == nil {
		panic("dfhackrpc: nil request")
	}
	if response == nil {
		panic("dfhackrpc: nil response")
	}

	requestType := reflect.TypeOf(request)
	responseType := reflect.TypeOf(response)

	if info, ok := c.methods[[2]string{plugin, method}]; ok {
		if requestType != info.request {
			panic("dfhackrpc: request type mismatch for method " + plugin + "::" + method + ": " + requestType.String() + " != " + info.request.String())
		}
		if responseType != info.response {
			panic("dfhackrpc: response type mismatch for method " + plugin + "::" + method + ": " + responseType.String() + " != " + info.response.String())
		}
		return info.id, OK, nil
	}

	var bindRequest dfproto_core.CoreBindRequest
	if plugin != "" {
		bindRequest.Plugin = proto.String(plugin)
	}
	bindRequest.Method = proto.String(method)
	bindRequest.InputMsg = proto.String(proto.MessageName(request))
	bindRequest.OutputMsg = proto.String(proto.MessageName(response))

	var bindResponse dfproto_core.CoreBindReply

	if rv, err := c.call("", bindMethod, &bindRequest, &bindResponse); err != nil {
		return 0, LinkFailure, err
	} else if rv != OK {
		return 0, rv, nil
	}

	id := int16(bindResponse.GetAssignedId())
	c.methods[[2]string{plugin, method}] = &methodInfo{
		request:  requestType,
		response: responseType,
		id:       id,
	}

	return id, OK, nil
}

func (c *Client) call(plugin, method string, request, response proto.Message) (CommandResult, error) {
	id, rv, err := c.bind(plugin, method, request, response)
	if rv != OK {
		return rv, err
	}

	if err := c.writeRequest(id, request); err != nil {
		return LinkFailure, err
	}

	return c.readResponse(response)
}
