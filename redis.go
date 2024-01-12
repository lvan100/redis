package redis

type ClientOption func(*Client)

func CustomStringOps(fn func(StringOps) StringOps) ClientOption {
	return func(c *Client) {
		c.stringOps = fn(c.stringOps)
	}
}

type Client struct {
	driver    Driver
	stringOps StringOps
}

func NewClient(driver Driver, opts ...ClientOption) (*Client, error) {
	c := &Client{
		driver:    driver,
		stringOps: &stringOps{driver},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c, nil
}

func (c *Client) StringOps() StringOps { return c.stringOps }
