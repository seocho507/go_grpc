package cmd

import (
	"fmt"
	"go_grpc/config"
	"go_grpc/grpc/client"
	"go_grpc/network"
	"go_grpc/repository"
	"go_grpc/service"
)

type App struct {
	config *config.Config

	gRPCClient *client.GRPCClient
	service    *service.Service
	network    *network.Network
	repository *repository.Repository
}

func NewApp(config *config.Config) (*App, error) {
	repo, err := repository.NewRepository(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create repository: %w", err)
	}

	svc, err := service.NewService(config, repo)
	if err != nil {
		return nil, fmt.Errorf("failed to create service: %w", err)
	}

	cli := client.NewGRPCClient(config)

	net, err := network.NewNetwork(config, svc, cli)
	if err != nil {
		return nil, fmt.Errorf("failed to create network: %w", err)
	}

	app := &App{
		config:     config,
		gRPCClient: cli,
		service:    svc,
		network:    net,
		repository: repo,
	}

	app.network.StartServer()
	return app, nil
}
