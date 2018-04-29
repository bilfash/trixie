package handlers

import (
	"encoding/json"
	"github.com/bilfash/trixie/config"
	"github.com/bilfash/trixie/domains"
	"github.com/bilfash/trixie/interfaces/api/requests"
	"github.com/bilfash/trixie/interfaces/api/responses"
	"github.com/bilfash/trixie/interfaces/pubsub"
	"github.com/qiangxue/fasthttp-routing"
)

type IClientHandler interface {
	SendMessageToTopic(topic string, message []byte) error
}

type ClientHandler struct {
	config        config.Configuration
	kafkaProducer IClientHandler
}

func NewClientHandler(config config.Configuration, kafkaProducer IClientHandler) ClientHandler {
	return ClientHandler{config, kafkaProducer}
}

func (t *ClientHandler) ClientApiPostHandler(c *routing.Context) error {
	req := &requests.ClientPost{}
	err := json.Unmarshal(c.PostBody(), &req)

	if err != nil {
		c.Response.SetStatusCode(responses.GetJsonRequestNotValidInstance().HttpStatus)
		c.Response.SetBody(responses.GetJsonRequestNotValidInstance().Error.GetError())
		return err
	}

	if req.Code == "" {
		c.Response.SetStatusCode(responses.GetBadRequestMissingParameterCodeInstance().HttpStatus)
		c.Response.SetBody(responses.GetBadRequestMissingParameterCodeInstance().Error.GetError())
		return err
	}

	if req.Name == "" {
		c.Response.SetStatusCode(responses.GetBadRequestMissingParameterNameInstance().HttpStatus)
		c.Response.SetBody(responses.GetBadRequestMissingParameterNameInstance().Error.GetError())
		return err
	}

	if req.IsActive == nil {
		c.Response.SetStatusCode(responses.GetBadRequestMissingParameterIsActiveInstance().HttpStatus)
		c.Response.SetBody(responses.GetBadRequestMissingParameterIsActiveInstance().Error.GetError())
		return err
	}

	clientMessage := pubsub.ClientMessage{
		Client: *domains.NewClient(req.Name, req.Code, *req.IsActive),
		ReqId:  "THISISUUID",
	}
	jsonMessage, _ := json.Marshal(clientMessage)
	err = t.kafkaProducer.SendMessageToTopic(t.config.KafkaProducerConfig.Topic.ClientCreateRequestedTopic, jsonMessage)

	if err != nil {
		c.Response.SetStatusCode(responses.GetUnknownErrorInstance().HttpStatus)
		c.Response.SetBody(responses.GetUnknownErrorInstance().Error.GetError())
		return err
	}

	c.Response.SetStatusCode(responses.GetAcceptedOkInstance().HttpStatus)
	return nil
}
