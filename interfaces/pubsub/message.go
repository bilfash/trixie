package pubsub

import "github.com/bilfash/trixie/domains"

type ClientMessage struct {
	Client domains.Client
	ReqId  string
}
