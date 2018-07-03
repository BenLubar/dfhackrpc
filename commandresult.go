package dfhackrpc

// CommandResult is a DFHack command result.
type CommandResult int

const (
	// LinkFailure is a failure due to an I/O or protocol error.
	LinkFailure CommandResult = -3
	// NeedsConsole is a failure due to a command requiring an interactive
	// context.
	NeedsConsole CommandResult = -2
	// NotImplemented is a failure due to a command not being implemented
	// or a plugin not being loaded.
	NotImplemented CommandResult = -1
	// OK is success.
	OK CommandResult = 0
	// Failure is a generic failure code.
	Failure CommandResult = 1
	// WrongUsage is a failure due to passing the wrong arguments or using
	// a command at a time when it is not valid.
	WrongUsage CommandResult = 2
	// NotFound is a failure due to something being missing, usually a
	// method or command.
	NotFound CommandResult = 3
)

// String implements fmt.Stringer
func (rv CommandResult) String() string {
	switch rv {
	case LinkFailure:
		return "RPC call failed due to I/O or protocol error"
	case NeedsConsole:
		return "Attempt to call interactive command without console"
	case NotImplemented:
		return "Command not implemented, or plugin not loaded"
	case OK:
		return "Success"
	case Failure:
		return "Failure"
	case WrongUsage:
		return "Wrong arguments or UI state"
	case NotFound:
		return "Target object not found"
	}
	return "?"
}

// Err is a convenience function that returns an error with the value from
// rv.String(), or nil if rv is OK.
func (rv CommandResult) Err() error {
	if rv == OK {
		return nil
	}

	return commandResultError(rv)
}

type commandResultError CommandResult

func (rv commandResultError) Error() string {
	return CommandResult(rv).String()
}
