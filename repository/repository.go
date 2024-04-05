package repository

import "go_grpc/config"

type Repository struct {
	config *config.Config
}

func NewRepository(config *config.Config) (*Repository, error) {
	return &Repository{config: config}, nil
}
