package applicator

import (
	"github.com/vshigimoto/GoAuthJWTService/internal/config"
	"github.com/vshigimoto/GoAuthJWTService/internal/server"
	"go.uber.org/zap"
)

type Applicator struct {
	l   *zap.Logger
	cfg *config.Config
}

func New() *Applicator {
	return &Applicator{}
}

func (a *Applicator) Start() error {

	server := server.New(a.cfg)
	return nil
}
