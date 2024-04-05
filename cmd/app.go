package cmd

import (
	"fmt"
	"go_grpc/config"
	"go_grpc/network"
	"go_grpc/repository"
	"go_grpc/service"
)

type App struct {
	config     *config.Config
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

	net, err := network.NewNetwork(config, svc)
	if err != nil {
		return nil, fmt.Errorf("failed to create network: %w", err)
	}

	return &App{
		config:     config,
		service:    svc,
		network:    net,
		repository: repo,
	}, nil
}
