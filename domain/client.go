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
	if c.id == otherClient.id && c.name == otherClient.name && c.code == otherClient.code && c.isActive == otherClient.isActive {
		return true
	}
	return false
}
