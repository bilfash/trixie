package pubsub

import (
	"fmt"
)

type CreateClientService struct {
}

func NewCreateClientService() CreateClientService {
	return CreateClientService{}
}

func (ccs *CreateClientService) Execute(eventData ClientMessage) {
	fmt.Println("event data", eventData.ReqId, eventData.Client)
}
