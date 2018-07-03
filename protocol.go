package dfhackrpc

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"

	"github.com/BenLubar/dfhackrpc/dfproto"
	"github.com/golang/protobuf/proto"
)

// version 1
const clientMagic = "DFHack?\n\x01\x00\x00\x00"
const serverMagic = "DFHack!\n\x01\x00\x00\x00"

const maxMessageSize = 64 * 1048576

const (
	rpcReplyResult = -1
	rpcReplyFail   = -2
	rpcReplyText   = -3
	rpcRequestQuit = -4
)

type rpcMessageHeader struct {
	ID   int16
	Size int32
}

func (c *Client) handshake() (err error) {
	defer func() {
		if err != nil {
			_ = c.conn.Close()
			c.conn = nil
		}
	}()

	if n, err := c.conn.Write([]byte(clientMagic)); err != nil {
		return err
	} else if n != len(clientMagic) {
		return io.ErrShortWrite
	}

	var response [len(serverMagic)]byte
	if _, err := io.ReadFull(c.conn, response[:]); err == io.EOF {
		return io.ErrUnexpectedEOF
	} else if err != nil {
		return err
	}

	if !bytes.Equal(response[:], []byte(serverMagic)) {
		return errors.New("dfhackrpc: invalid handshake response from server")
	}

	c.resetMethods()
	c.bad = false
	return nil
}

func (c *Client) writeHeader(id int16, size int) error {
	if size > maxMessageSize {
		return errors.New("dfhackrpc: message is too large")
	}

	return binary.Write(c.conn, binary.LittleEndian, &rpcMessageHeader{
		ID:   id,
		Size: int32(size),
	})
}

func (c *Client) writeRequest(id int16, v proto.Message) error {
	b, err := proto.Marshal(v)
	if err != nil {
		return err
	}

	if err = c.writeHeader(id, len(b)); err != nil {
		return err
	}

	if n, err := c.conn.Write(b); err != nil {
		c.bad = true
		return err
	} else if n != len(b) {
		c.bad = true
		return io.ErrShortWrite
	}

	return nil
}

func (c *Client) readResponse(v proto.Message) (CommandResult, error) {
	for {
		var header rpcMessageHeader
		if err := binary.Read(c.conn, binary.LittleEndian, &header); err != nil {
			c.bad = true
			return LinkFailure, err
		}

		switch header.ID {
		case rpcReplyText:
			var text dfproto.CoreTextNotification
			if err := c.readMessage(header.Size, &text); err != nil {
				return LinkFailure, err
			}
			for _, fragment := range text.GetFragments() {
				c.outputTextFragment(fragment)
			}
		case rpcReplyFail:
			rv := CommandResult(header.Size)
			if rv == OK {
				rv = Failure
			}
			return rv, nil
		case rpcReplyResult:
			if err := c.readMessage(header.Size, v); err != nil {
				return LinkFailure, err
			}
			return OK, nil
		}
	}
}

func (c *Client) readMessage(size int32, v proto.Message) error {
	if size > maxMessageSize {
		c.bad = true
		return errors.New("dfhackrpc: response message exceeds maximum size")
	}

	b := make([]byte, size)
	if _, err := io.ReadFull(c.conn, b); err == io.EOF {
		c.bad = true
		return io.ErrUnexpectedEOF
	} else if err != nil {
		c.bad = true
		return err
	}

	return proto.Unmarshal(b, v)
}
