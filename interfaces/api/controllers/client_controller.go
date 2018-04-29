package controllers

import (
	"encoding/json"
	"github.com/bilfash/trixie/domains"
	"github.com/bilfash/trixie/interfaces/api/models/requests"
	"github.com/bilfash/trixie/interfaces/api/models/responses"
	"github.com/qiangxue/fasthttp-routing"
)

type IClientController interface {
	SendMessage(message []byte) error
}

type ClientController struct {
	kafkaProducer IClientController
}

func NewClientController(kafkaProducer IClientController) ClientController {
	return ClientController{kafkaProducer}
}

func (t *ClientController) ClientApiPostHandler(c *routing.Context) error {
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

	clientObj := domains.NewClient(req.Name, req.Code, *req.IsActive)
	jsonMessage, err := json.Marshal(clientObj)
	err = t.kafkaProducer.SendMessage(jsonMessage)

	if err != nil {
		c.Response.SetStatusCode(responses.GetUnknownErrorInstance().HttpStatus)
		c.Response.SetBody(responses.GetUnknownErrorInstance().Error.GetError())
		return err
	}

	c.Response.SetStatusCode(responses.GetAcceptedOkInstance().HttpStatus)
	return nil
}
