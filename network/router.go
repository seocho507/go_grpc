package network

import (
	"go_grpc/config"
	"go_grpc/service"
)

type Network struct {
	config  *config.Config
	service *service.Service
}

func NewNetwork(config *config.Config, service *service.Service) (*Network, error) {
	return &Network{
		config:  config,
		service: service,
	}, nil
}
