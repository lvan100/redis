package redis

import (
	"bytes"
	"sync"
)

var bytesPool sync.Pool

func getBytes() *bytes.Buffer {
	if v := bytesPool.Get(); v != nil {
		buf := v.(*bytes.Buffer)
		buf.Reset()
		return buf
	}
	return bytes.NewBuffer(nil)
}

func putBytes(v *bytes.Buffer) {
	bytesPool.Put(v)
}

type Client struct {
	driver    Driver
	stringCmd StringCmd
}

type ClientOption func(*Client)

func CustomStringCmd(fn func(StringCmd) StringCmd) ClientOption {
	return func(c *Client) {
		c.stringCmd = fn(c.stringCmd)
	}
}

func NewClient(driver Driver, opts ...ClientOption) (*Client, error) {
	c := &Client{driver: driver}
	c.stringCmd = &stringCmd{driver: driver}
	for _, opt := range opts {
		opt(c)
	}
	return c, nil
}

func (c *Client) StringCmd() StringCmd {
	return c.stringCmd
}
