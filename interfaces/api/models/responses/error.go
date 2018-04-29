package responses

import "fmt"

type errorResp struct {
	Code    int
	Message string
	Errors  errors
}

type errors struct {
	Domain  string
	Reason  string
	Message string
}

func (e *errorResp) GetError() []byte {
	str := fmt.Sprintf(`{"error": { "errors": {"domain": "%s","reason": "%s","message": "%s"},"code": %d,"message": "%s"}}`,
		e.Errors.Domain, e.Errors.Reason, e.Errors.Message, e.Code, e.Message)
	return []byte(str)
}
