package network

import (
	"github.com/gin-gonic/gin"
	"go_grpc/config"
	"go_grpc/service"
	"os"
)

type Network struct {
	config  *config.Config
	service *service.Service
	engin   *gin.Engine
}

func NewNetwork(config *config.Config, service *service.Service) (*Network, error) {
	return &Network{
		config:  config,
		service: service,
		engin:   gin.New(),
	}, nil
}

func (n *Network) StartServer() {
	err := n.engin.Run(":8080")
	if err != nil {
		os.Exit(100)
	}
}
