package redis

type ClientOption func(*Client)

func CustomStringOps(fn func(StringOps) StringOps) ClientOption {
	return func(c *Client) {
		c.stringOps = fn(c.stringOps)
	}
}

func CustomBitmapOps(fn func(ops BitmapOps) BitmapOps) ClientOption {
	return func(c *Client) {
		c.bitmapOps = fn(c.bitmapOps)
	}
}

func CustomHashOps(fn func(HashOps) HashOps) ClientOption {
	return func(c *Client) {
		c.hashOps = fn(c.hashOps)
	}
}

func CustomKeyOps(fn func(KeyOps) KeyOps) ClientOption {
	return func(c *Client) {
		c.keyOps = fn(c.keyOps)
	}
}

func CustomListOps(fn func(ListOps) ListOps) ClientOption {
	return func(c *Client) {
		c.listOps = fn(c.listOps)
	}
}

func CustomSetOps(fn func(SetOps) SetOps) ClientOption {
	return func(c *Client) {
		c.setOps = fn(c.setOps)
	}
}

func CustomSortedSetOps(fn func(SortedSetOps) SortedSetOps) ClientOption {
	return func(c *Client) {
		c.sortedSetOps = fn(c.sortedSetOps)
	}
}

type Client struct {
	stringOps    StringOps
	bitmapOps    BitmapOps
	hashOps      HashOps
	keyOps       KeyOps
	listOps      ListOps
	setOps       SetOps
	sortedSetOps SortedSetOps
}

func NewClient(driver Driver, opts ...ClientOption) (*Client, error) {
	c := &Client{
		stringOps:    &stringOps{driver},
		bitmapOps:    &bitmapOps{driver},
		hashOps:      &hashOps{driver},
		keyOps:       &keyOps{driver},
		listOps:      &listOps{driver},
		setOps:       &setOps{driver},
		sortedSetOps: &sortedSetOps{driver},
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

func (c *Client) SetOps() SetOps { return c.setOps }

func (c *Client) SortedSetOps() SortedSetOps { return c.sortedSetOps }
