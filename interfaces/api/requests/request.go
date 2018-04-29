package requests

type ClientPost struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	IsActive *bool  `json:"isActive"`
}
