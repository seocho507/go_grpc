package server

import (
	"context"
	"fmt"
	"go_grpc/config"
	"go_grpc/grpc/paseto"
	authpb "go_grpc/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	okResponse    = authpb.ResponseType_OK
	errorResponse = authpb.ResponseType_ERROR
)

type GRPCServer struct {
	authpb.UnimplementedAuthServiceServer
	PasetoMaker    *paseto.PasetoMaker
	tokenVerifiers map[string]*authpb.Auth
}

func NewGRPCServer(cfg *config.Config) (*GRPCServer, error) {
	listen, err := net.Listen("tcp", cfg.GRPC.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}

	server := grpc.NewServer()
	grpcServer := &GRPCServer{
		PasetoMaker:    paseto.NewPasetoMaker(cfg),
		tokenVerifiers: make(map[string]*authpb.Auth),
	}

	authpb.RegisterAuthServiceServer(server, grpcServer)
	reflection.Register(server)

	go func() {
		log.Println("Starting gRPC server...")
		log.Println("Listening on", cfg.GRPC.URL)
		if err := server.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	return grpcServer, nil
}

func (s *GRPCServer) CreateToken(_ context.Context, req *authpb.CreateTokenRequest) (*authpb.CreateTokenResponse, error) {
	data := req.GetAuth()
	token := data.GetToken()
	s.tokenVerifiers[token] = data

	return &authpb.CreateTokenResponse{
		Type: okResponse,
		Auth: data,
	}, nil
}

func (s *GRPCServer) ValidateToken(_ context.Context, req *authpb.ValidateTokenRequest) (*authpb.ValidateTokenResponse, error) {
	token := req.GetToken()
	data, err := s.getTokenData(token)
	if err != nil {
		return &authpb.ValidateTokenResponse{
			Type: errorResponse,
		}, nil
	}

	return &authpb.ValidateTokenResponse{
		Type: okResponse,
		Auth: data,
	}, nil
}

func (s *GRPCServer) getTokenData(token string) (*authpb.Auth, error) {
	data, ok := s.tokenVerifiers[token]
	if !ok {
		return nil, fmt.Errorf("token not found: %s", token)
	}

	if data.GetToken() != token {
		return nil, fmt.Errorf("token mismatch: expected %s, got %s", token, data.GetToken())
	}

	return data, nil
}
