package responses

import "sync"

type badRequestMissingParameterName struct {
	HttpStatus int
	Error      errorResp
}

var badRequestMissingParameterNameInstance *badRequestMissingParameterName
var onceBadRequestMissingParameterName sync.Once

func GetBadRequestMissingParameterNameInstance() *badRequestMissingParameterName {
	onceBadRequestMissingParameterName.Do(func() {
		badRequestMissingParameterNameInstance = &badRequestMissingParameterName{
			400,
			errorResp{
				400,
				"missing parameter name",
				errors{
					"client management",
					"error when try to unmarshal json request, missing paramater name",
					"missing parameter name",
				},
			}}
	})
	return badRequestMissingParameterNameInstance
}
