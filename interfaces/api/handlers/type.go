package handlers

type IClientHandler interface {
	SendMessageToTopic(topic string, message []byte) error
}

type IResponseHTTP interface {
	SetStatusCode(statusCode int)
	SetBody(body []byte)
}
