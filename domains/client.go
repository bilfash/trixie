package domains

type ClientRepo interface {
	Insert(client Client) error
}

type Client struct {
	Code     string
	name     string
	isActive bool
}

func NewClient(name string, code string, isActive bool) *Client {
	return &Client{code, name, isActive}
}

func (c *Client) EqualTo(otherClient Client) bool {
	if c.name == otherClient.name && c.Code == otherClient.Code && c.isActive == otherClient.isActive {
		return true
	}
	return false
}
