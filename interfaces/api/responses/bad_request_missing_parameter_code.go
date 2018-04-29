package responses

import "sync"

type badRequestMissingParameterCode struct {
	HttpStatus int
	Error      errorResp
}

var badRequestMissingParameterCodeInstance *badRequestMissingParameterCode
var onceBadRequestMissingParameterCode sync.Once

func GetBadRequestMissingParameterCodeInstance() *badRequestMissingParameterCode {
	onceBadRequestMissingParameterCode.Do(func() {
		badRequestMissingParameterCodeInstance = &badRequestMissingParameterCode{
			400,
			errorResp{
				400,
				"missing parameter code",
				errors{
					"client management",
					"error when try to unmarshal json request, missing paramater code",
					"missing parameter code",
				},
			}}
	})
	return badRequestMissingParameterCodeInstance
}
