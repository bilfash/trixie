package responses

import "sync"

type unknownError struct {
	HttpStatus int
	Error      errorResp
}

var unknownErrorInstance *unknownError
var onceunknownError sync.Once

func GetUnknownErrorInstance() *unknownError {
	onceunknownError.Do(func() {
		unknownErrorInstance = &unknownError{
			500,
			errorResp{
				500,
				"unknown error",
				errors{
					"global",
					"unknown error",
					"please contact tech team",
				},
			}}
	})
	return unknownErrorInstance
}
