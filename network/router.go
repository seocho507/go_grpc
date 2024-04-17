package network

import (
	"github.com/gin-gonic/gin"
	"go_grpc/config"
	"go_grpc/grpc/client"
	"go_grpc/service"
	"os"
)

type Network struct {
	config *config.Config

	gRPCClient *client.GRPCClient
	service    *service.Service
	engin      *gin.Engine
}

func NewNetwork(config *config.Config, service *service.Service, cli *client.GRPCClient) (*Network, error) {

	r := &Network{
		config:     config,
		gRPCClient: cli,
		service:    service,
		engin:      gin.New(),
	}

	// 1. API for CreateToken
	r.engin.POST("/login", r.login)

	// 2. API for ValidateToken
	r.engin.GET("/validate", r.verify)

	return r, nil
}

func (n *Network) StartServer() {
	err := n.engin.Run(":8080")
	if err != nil {
		os.Exit(100)
	}
}
