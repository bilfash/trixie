package domain

type Client struct {
	id       int64
	name     string
	code     string
	isActive bool
}

func NewClient(id int64, name string, code string, isActive bool) *Client {
	return &Client{id, name, code, isActive}
}

func (c *Client) EqualTo(otherClient Client) bool {
	return true
}
