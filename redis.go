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
	bitmapOps BitmapOps
	hashOps   HashOps
	keyOps    KeyOps
	listOps   ListOps
}

func NewClient(driver Driver, opts ...ClientOption) (*Client, error) {
	c := &Client{
		driver:    driver,
		stringOps: &stringOps{driver},
		bitmapOps: &bitmapOps{driver},
		hashOps:   &hashOps{driver},
		keyOps:    &keyOps{driver},
		listOps:   &listOps{driver},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c, nil
}

func (c *Client) StringOps() StringOps { return c.stringOps }

func (c *Client) BitmapOps() BitmapOps { return c.bitmapOps }

func (c *Client) HashOps() HashOps { return c.hashOps }

func (c *Client) KeyOps() KeyOps { return c.keyOps }

func (c *Client) ListOps() ListOps { return c.listOps }
