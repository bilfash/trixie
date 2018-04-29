package responses

import "sync"

type acceptedOk struct {
	HttpStatus int
}

var acceptedOkInstance *acceptedOk
var onceacceptedOk sync.Once

func GetAcceptedOkInstance() *acceptedOk {
	onceacceptedOk.Do(func() {
		acceptedOkInstance = &acceptedOk{
			200,
		}
	})
	return acceptedOkInstance
}
