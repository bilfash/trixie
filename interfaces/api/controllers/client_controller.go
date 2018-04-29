package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bilfash/trixie/interfaces/api/models/requests"
	"github.com/bilfash/trixie/interfaces/api/models/responses"
)

type ClientController struct {
}

func NewClientController() ClientController {
	return ClientController{}
}

func (t *ClientController) ClientApiPostHandler(w http.ResponseWriter, request *http.Request) {
	bodyBytes, _ := ioutil.ReadAll(request.Body)
	req := &requests.ClientPost{}
	err := json.Unmarshal([]byte(bodyBytes), &req)

	if err != nil {
		w.WriteHeader(responses.GetJsonRequestNotValidInstance().HttpStatus)
		w.Write(responses.GetJsonRequestNotValidInstance().Error.GetError())
		return
	}

	if req.Code == "" {
		w.WriteHeader(responses.GetBadRequestMissingParameterCodeInstance().HttpStatus)
		w.Write(responses.GetBadRequestMissingParameterCodeInstance().Error.GetError())
		return
	}

	if req.Name == "" {
		w.WriteHeader(responses.GetBadRequestMissingParameterNameInstance().HttpStatus)
		w.Write(responses.GetBadRequestMissingParameterNameInstance().Error.GetError())
		return
	}

	if req.IsActive == nil {
		w.WriteHeader(responses.GetBadRequestMissingParameterIsActiveInstance().HttpStatus)
		w.Write(responses.GetBadRequestMissingParameterIsActiveInstance().Error.GetError())
		return
	}

	w.WriteHeader(http.StatusOK)
}
