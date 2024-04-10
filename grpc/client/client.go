package client

import (
	"context"
	"go_grpc/config"
	"go_grpc/grpc/paseto"
	auth "go_grpc/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type GRPCClient struct {
	client      *grpc.ClientConn
	authClient  auth.AuthServiceClient
	pasetoMaker *paseto.PasetoMaker
}

func NewGRPCClient(config *config.Config) *GRPCClient {
	client, err := grpc.Dial(config.GRPC.URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}
	return &GRPCClient{
		client:      client,
		authClient:  auth.NewAuthServiceClient(client),
		pasetoMaker: paseto.NewPasetoMaker(config),
	}
}

/*
service AuthService {
	rpc CreateToken(CreateTokenRequest) returns (CreateTokenResponse);
	rpc ValidateToken(ValidateTokenRequest) returns (ValidateTokenResponse);
}
*/

func (g *GRPCClient) CreateToken(name string) (*auth.CreateTokenResponse, error) {
	now := time.Now()
	expires := now.Add(time.Duration(24) * time.Hour)

	req := &auth.CreateTokenRequest{
		Auth: &auth.Auth{
			Name:    name,
			Token:   g.pasetoMaker.CreateToken(name),
			Created: now.Unix(),
			Expires: expires.Unix(),
		},
	}

	res, err := g.authClient.CreateToken(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (g *GRPCClient) ValidateToken(token string) (*auth.ValidateTokenResponse, error) {
	req := &auth.ValidateTokenRequest{
		Token: token,
	}

	res, err := g.authClient.ValidateToken(context.Background(), req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
