package dfhackrpc

import (
	"bytes"
	"encoding/binary"
	"io"
	"net"

	"github.com/BenLubar/dfhackrpc/dfproto/core"
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

func writeFull(w io.Writer, b []byte) error {
	n, err := w.Write(b)

	if err == nil && n != len(b) {
		err = io.ErrShortWrite
	}

	return err
}

func readFull(r io.Reader, b []byte) error {
	_, err := io.ReadFull(r, b)

	if err == io.EOF {
		err = io.ErrUnexpectedEOF
	}

	return err
}

func (c *Client) handshake(conn net.Conn) error {
	if err := writeFull(conn, []byte(clientMagic)); err != nil {
		_ = conn.Close()
		return err
	}

	var response [len(serverMagic)]byte
	if err := readFull(conn, response[:]); err != nil {
		_ = conn.Close()
		return err
	}

	if !bytes.Equal(response[:], []byte(serverMagic)) {
		_ = conn.Close()
		return errBadHandshake
	}

	c.resetMethods()
	c.bad = false
	c.conn = conn
	return nil
}

func (c *Client) writeHeader(id int16, size int) error {
	if size > maxMessageSize {
		return errRequestTooLarge
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

	if err := writeFull(c.conn, b); err != nil {
		c.bad = true
		return err
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
			var text dfproto_core.CoreTextNotification
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
		return errResponseTooLarge
	}

	b := make([]byte, size)
	if err := readFull(c.conn, b); err != nil {
		c.bad = true
		return err
	}

	err := proto.Unmarshal(b, v)
	if err != nil {
		c.bad = true
	}

	return err
}
