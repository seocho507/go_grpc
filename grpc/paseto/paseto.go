package paseto

import (
	"github.com/o1egl/paseto"
	"go_grpc/config"
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

func (p *PasetoMaker) CreateToken(data string) (string, error) {
	return p.Pt.Encrypt(p.Key, []byte(data), nil)
}

func (p *PasetoMaker) VerifyToken(token string) (string, error) {
	var data string
	err := p.Pt.Decrypt(token, p.Key, []byte{}, &data)
	if err != nil {
		return "", err
	}
	return data, nil
}
