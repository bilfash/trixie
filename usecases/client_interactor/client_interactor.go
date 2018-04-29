package client_interactor

import "github.com/bilfash/trixie/domains"

type ClientInteractor struct {
	clientRepo domains.ClientRepository
}

func NewClientInteractor(repository domains.ClientRepository) *ClientInteractor {
	return &ClientInteractor{repository}
}

func (ci *ClientInteractor) Store(client domains.Client) error {
	err := ci.clientRepo.Store(client)
	return err
}
