package dfhackrpc

import "errors"

var (
	errAlreadyConnected = errors.New("dfhackrpc: client is already connected")
	errNotConnected     = errors.New("dfhackrpc: not connected")
	errBadHandshake     = errors.New("dfhackrpc: invalid handshake response from server")
	errRequestTooLarge  = errors.New("dfhackrpc: request message exceeds maximum size")
	errResponseTooLarge = errors.New("dfhackrpc: response message exceeds maximum size")
)
