package redis

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
