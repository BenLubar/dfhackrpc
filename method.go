package dfhackrpc

import (
	"reflect"

	"github.com/BenLubar/dfhackrpc/dfproto"
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
	c.methods = map[string]*methodInfo{
		bindMethod: {
			request:  reflect.TypeOf(&dfproto.CoreBindRequest{}),
			response: reflect.TypeOf(&dfproto.CoreBindReply{}),
			id:       0,
		},
		runCommand: {
			request:  reflect.TypeOf(&dfproto.CoreRunCommandRequest{}),
			response: reflect.TypeOf(&dfproto.EmptyMessage{}),
			id:       1,
		},
	}
}

// Call runs a RPC method. It is an error to call a method with the wrong
// request or response type. If request or response are nil or their types
// do not match the types used the first time the method was called, Call
// will panic.
func (c *Client) Call(method string, request, response proto.Message) (CommandResult, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if err := c.checkReconnect(); err != nil {
		return LinkFailure, err
	}

	return c.call(method, request, response)
}

func (c *Client) bind(method string, request, response proto.Message) (int16, CommandResult, error) {
	if request == nil {
		panic("dfhackrpc: nil request")
	}
	if response == nil {
		panic("dfhackrpc: nil response")
	}

	requestType := reflect.TypeOf(request)
	responseType := reflect.TypeOf(response)

	if info, ok := c.methods[method]; ok {
		if requestType != info.request {
			panic("dfhackrpc: request type mismatch for method " + method + ": " + requestType.String() + " != " + info.request.String())
		}
		if responseType != info.response {
			panic("dfhackrpc: response type mismatch for method " + method + ": " + responseType.String() + " != " + info.response.String())
		}
		return info.id, OK, nil
	}

	var bindRequest dfproto.CoreBindRequest
	bindRequest.Method = proto.String(method)
	bindRequest.InputMsg = proto.String(proto.MessageName(request))
	bindRequest.OutputMsg = proto.String(proto.MessageName(response))

	var bindResponse dfproto.CoreBindReply

	if rv, err := c.call(bindMethod, &bindRequest, &bindResponse); err != nil {
		return 0, LinkFailure, err
	} else if rv != OK {
		return 0, rv, nil
	}

	id := int16(bindResponse.GetAssignedId())
	c.methods[method] = &methodInfo{
		request:  requestType,
		response: responseType,
		id:       id,
	}

	return id, OK, nil
}

func (c *Client) call(method string, request, response proto.Message) (CommandResult, error) {
	id, rv, err := c.bind(method, request, response)
	if rv != OK {
		return rv, err
	}

	if err := c.writeRequest(id, request); err != nil {
		return LinkFailure, err
	}

	return c.readResponse(response)
}

// RunCommand is a convenience wrapper equivalent to calling Call with the
// method "RunCommand", a *dfproto.CoreRunCommandRequest request, and a
// *dfproto.EmptyMessage response.
func (c *Client) RunCommand(command string, args ...string) (CommandResult, error) {
	var request dfproto.CoreRunCommandRequest
	request.Command = proto.String(command)
	request.Arguments = args

	var response dfproto.EmptyMessage

	return c.Call(runCommand, &request, &response)
}
