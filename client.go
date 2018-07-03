// Package dfhackrpc provides a DFHack RPC client.
package dfhackrpc // import "github.com/BenLubar/dfhackrpc"

import (
	"io"
	"net"
	"os"
	"runtime"
	"strconv"
	"sync"

	"github.com/mattn/go-colorable"
)

// Client is a DFHack RPC client. Use NewClient to create an instance of this
// type. All methods on Client can be called safely from multiple goroutines.
//
// Client will automatically reconnect if a protocol error was encountered on
// the previous method call.
type Client struct {
	out     io.Writer
	addr    string
	conn    net.Conn
	bad     bool
	methods map[string]*methodInfo
	lock    sync.Mutex
}

// NewClient constructs a client that writes output to w. The output will use
// colors if w is a terminal (TTY).
func NewClient(w io.Writer) *Client {
	c := &Client{out: wrapWriter(w)}
	runtime.SetFinalizer(c, (*Client).Close)
	return c
}

// NewClientColor constructs a client that writes output to w. The output will
// use colors if and only if color is set to true.
func NewClientColor(w io.Writer, color bool) *Client {
	var c *Client
	if !color {
		c = &Client{out: colorable.NewNonColorable(w)}
	} else if f, ok := w.(*os.File); ok {
		c = &Client{out: colorable.NewColorable(f)}
	} else {
		c = &Client{out: w}
	}
	runtime.SetFinalizer(c, (*Client).Close)
	return c
}

// Connect connects to localhost on the port given by the DFHACK_PORT
// environment variable, or 5000 if it is not set to a positive integer.
func (c *Client) Connect() error {
	port, err := strconv.Atoi(os.Getenv("DFHACK_PORT"))
	if err != nil || port <= 0 {
		port = 5000
	}
	return c.ConnectAddr("localhost:" + strconv.Itoa(port))
}

// ConnectAddr connects to a specified TCP address. If the Client is already
// connected, an error is returned. If any error occurs while connecting, the
// client will be in a valid disconnected state when this method returns.
func (c *Client) ConnectAddr(addr string) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.conn != nil {
		return errAlreadyConnected
	}

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	c.addr = addr
	return c.handshake(conn)
}

// Close disconnects the client. When this method returns, the client will be
// in a valid disconnected state regardless of the error returned.
func (c *Client) Close() error {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.conn == nil {
		return nil
	}

	var err1 error
	if !c.bad {
		err1 = c.writeHeader(rpcRequestQuit, 0)
	}
	err2 := c.conn.Close()

	c.bad = false
	c.addr = ""
	c.conn = nil
	c.methods = nil

	if err1 != nil {
		return err1
	}
	return err2
}

func (c *Client) checkReconnect() error {
	if c.conn == nil {
		return errNotConnected
	}

	if !c.bad {
		return nil
	}

	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return err
	}

	oldConn := c.conn
	if err = c.handshake(conn); err == nil {
		_ = oldConn.Close()
	}

	return err
}
