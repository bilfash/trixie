package responses

import "sync"

type badRequestMissingParameterIsActive struct {
	HttpStatus int
	Error      errorResp
}

var badRequestMissingParameterIsActiveInstance *badRequestMissingParameterIsActive
var onceBadRequestMissingParameterIsActive sync.Once

func GetBadRequestMissingParameterIsActiveInstance() *badRequestMissingParameterIsActive {
	onceBadRequestMissingParameterIsActive.Do(func() {
		badRequestMissingParameterIsActiveInstance = &badRequestMissingParameterIsActive{
			400,
			errorResp{
				400,
				"missing parameter isActive",
				errors{
					"client management",
					"error when try to unmarshal json request, missing paramater isActive",
					"missing parameter isActive",
				},
			}}
	})
	return badRequestMissingParameterIsActiveInstance
}
