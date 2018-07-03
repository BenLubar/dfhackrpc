package dfhackrpc

import (
	"github.com/BenLubar/dfhackrpc/dfproto/core"
	"github.com/golang/protobuf/proto"
)

// RunCommand is a convenience wrapper around the RunCommand RPC method.
func (c *Client) RunCommand(command string, args ...string) (CommandResult, error) {
	var request dfproto_core.CoreRunCommandRequest
	request.Command = proto.String(command)
	request.Arguments = args

	var response dfproto_core.EmptyMessage

	return c.Call(runCommand, &request, &response)
}

// Lock is a convenience wrapper around the CoreSuspend RPC method.
// It implements the sync.Locker interface and panics if an error occurs.
func (c *Client) Lock() {
	c.callSuspend("CoreSuspend")
}

// Unlock is a convenience wrapper around the CoreResume RPC method.
// It implements the sync.Locker interface and panics if an error occurs.
func (c *Client) Unlock() {
	c.callSuspend("CoreResume")
}

func (c *Client) callSuspend(method string) {
	var request dfproto_core.EmptyMessage

	var response dfproto_core.IntMessage

	rv, err := c.Call(method, &request, &response)
	if err == nil {
		err = rv.Err()
	}

	if err != nil {
		panic(err)
	}
}

// RunLua is a convenience wrapper around the RunLua RPC method.
func (c *Client) RunLua(module, function string, args ...string) ([]string, CommandResult, error) {
	var request dfproto_core.CoreRunLuaRequest
	request.Module = proto.String(module)
	request.Function = proto.String(function)
	request.Arguments = args

	var response dfproto_core.StringListMessage

	rv, err := c.Call("RunLua", &request, &response)
	return response.GetValue(), rv, err
}
