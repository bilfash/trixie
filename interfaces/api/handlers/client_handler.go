package handlers

import (
	"encoding/json"
	"github.com/bilfash/trixie/config"
	"github.com/bilfash/trixie/domains"
	"github.com/bilfash/trixie/interfaces/api/requests"
	"github.com/bilfash/trixie/interfaces/api/responses"
	"github.com/bilfash/trixie/interfaces/pubsub"
)

type ClientHandler struct {
	config        config.Configuration
	kafkaProducer IClientHandler
}

func NewClientHandler(config config.Configuration, kafkaProducer IClientHandler) ClientHandler {
	return ClientHandler{config, kafkaProducer}
}

func (t *ClientHandler) ClientApiPostHandler(requestBody []byte, responseHTTP IResponseHTTP) error {
	req := &requests.ClientPost{}
	err := json.Unmarshal(requestBody, &req)

	if err != nil {
		responseHTTP.SetStatusCode(responses.GetJsonRequestNotValidInstance().HttpStatus)
		responseHTTP.SetBody(responses.GetJsonRequestNotValidInstance().Error.GetError())
		return err
	}

	if req.Code == "" {
		responseHTTP.SetStatusCode(responses.GetBadRequestMissingParameterCodeInstance().HttpStatus)
		responseHTTP.SetBody(responses.GetBadRequestMissingParameterCodeInstance().Error.GetError())
		return err
	}

	if req.Name == "" {
		responseHTTP.SetStatusCode(responses.GetBadRequestMissingParameterNameInstance().HttpStatus)
		responseHTTP.SetBody(responses.GetBadRequestMissingParameterNameInstance().Error.GetError())
		return err
	}

	if req.IsActive == nil {
		responseHTTP.SetStatusCode(responses.GetBadRequestMissingParameterIsActiveInstance().HttpStatus)
		responseHTTP.SetBody(responses.GetBadRequestMissingParameterIsActiveInstance().Error.GetError())
		return err
	}

	clientMessage := pubsub.ClientMessage{
		Client: *domains.NewClient(req.Name, req.Code, *req.IsActive),
		ReqId:  "THISISUUID",
	}
	jsonMessage, _ := json.Marshal(clientMessage)
	err = t.kafkaProducer.SendMessageToTopic(t.config.KafkaProducerConfig.Topic.ClientCreateRequestedTopic, jsonMessage)

	if err != nil {
		responseHTTP.SetStatusCode(responses.GetUnknownErrorInstance().HttpStatus)
		responseHTTP.SetBody(responses.GetUnknownErrorInstance().Error.GetError())
		return err
	}

	responseHTTP.SetStatusCode(responses.GetAcceptedOkInstance().HttpStatus)
	return nil
}
