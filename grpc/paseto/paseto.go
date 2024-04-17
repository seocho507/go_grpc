package paseto

import (
	"crypto/rand"
	"fmt"
	"github.com/o1egl/paseto"
	"go_grpc/config"
	auth "go_grpc/grpc/proto"
	"log"
)

type PasetoMaker struct {
	Pt  *paseto.V2
	Key []byte
}

func NewPasetoMaker(config *config.Config) *PasetoMaker {
	return &PasetoMaker{
		Pt:  paseto.NewV2(),
		Key: []byte(config.Paseto.Key),
	}
}

func (p *PasetoMaker) CreateToken(data string) string {
	randomBytes := make([]byte, 32)
	// 랜덤 바이트 배열 채우기
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal(err)
	}

	token, err := p.Pt.Encrypt(p.Key, []byte(data), randomBytes)
	if err != nil {
		log.Fatal(err)
	}
	return token
}

func (p *PasetoMaker) VerifyToken(token string) (*auth.Auth, error) {
	authData := new(auth.Auth)
	err := p.Pt.Decrypt(token, p.Key, authData, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt token: %w", err)
	}
	return authData, nil
}
