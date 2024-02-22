package handlers

import (
	"go.uber.org/zap"

	"github.com/vshigimoto/GoAuthJWTService/internal/services"
)

type Handlers struct {
	services *services.Services
	l        *zap.SugaredLogger
}

func New(services *services.Services, l *zap.SugaredLogger) *Handlers {
	return &Handlers{
		services: services,
		l:        l,
	}
}
