package redis

type ClientOption func(*Client)

func CustomStringCmd(fn func(StringCmd) StringCmd) ClientOption {
	return func(c *Client) {
		c.stringCmd = fn(c.stringCmd)
	}
}

type Client struct {
	driver    Driver
	stringCmd StringCmd
}

func NewClient(driver Driver, opts ...ClientOption) (*Client, error) {
	c := &Client{
		driver:    driver,
		stringCmd: &stringCmd{driver},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c, nil
}

func (c *Client) StringCmd() StringCmd { return c.stringCmd }
