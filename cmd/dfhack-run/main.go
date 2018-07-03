// Command dfhack-run is a port of the dfhack-run command in C++ to Go.
//
// It has two forms:
//
//     dfhack-run <command> [args...]
//
// Run a DFHack command as if it were typed into the console by the user.
//
//     dfhack-run --lua <module> <function> [args...]
//
// Run a Lua function from a specified module.
//
// It uses port 5000 by default, but this can be overridden using the
// DFHACK_PORT environment variable.
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/BenLubar/dfhackrpc"
	"github.com/BenLubar/dfhackrpc/dfproto"
	"github.com/golang/protobuf/proto"
)

func fatal(code int, args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(code)
	panic("unreachable")
}

func main() {
	if len(os.Args) <= 1 {
		fatal(2, "Usage: dfhack-run <command> [args...]")
	}

	// Connect to DFHack
	client := dfhackrpc.NewClient(os.Stdout)
	if err := client.Connect(); err != nil {
		fatal(2, err)
	}

	var rv dfhackrpc.CommandResult
	var err error

	if os.Args[1] == "--lua" {
		if len(os.Args) <= 3 {
			fatal(2, "Usage: dfhack-run --lua <module> <function> [args...]")
		}

		var request dfproto.CoreRunLuaRequest
		request.Module = proto.String(os.Args[2])
		request.Function = proto.String(os.Args[3])
		request.Arguments = os.Args[4:]

		var response dfproto.StringListMessage

		rv, err = client.Call("RunLua", &request, &response)
		if err != nil {
			fatal(1, err)
		}

		if rv == dfhackrpc.OK {
			fmt.Println(strings.Join(response.GetValue(), "\t"))
		}
	} else {
		// Call the command
		rv, err = client.RunCommand(os.Args[1], os.Args[2:]...)
		if err != nil {
			fatal(1, err)
		}
	}

	if rv != dfhackrpc.OK {
		fatal(1, rv)
	}
}
