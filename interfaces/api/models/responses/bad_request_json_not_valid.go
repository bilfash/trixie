package responses

import "sync"

type badRequestJsonRequestNotValid struct {
	HttpStatus int
	Error      errorResp
}

var badRequestJsonRequestNotValidInstance *badRequestJsonRequestNotValid
var oncebadRequestJsonRequestNotValid sync.Once

func GetJsonRequestNotValidInstance() *badRequestJsonRequestNotValid {
	oncebadRequestJsonRequestNotValid.Do(func() {
		badRequestJsonRequestNotValidInstance = &badRequestJsonRequestNotValid{
			400,
			errorResp{
				400,
				"json not valid",
				errors{
					"global",
					"error when try to unmarshal json request",
					"json not valid",
				},
			}}
	})
	return badRequestJsonRequestNotValidInstance
}
