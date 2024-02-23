package handlers

import (
	"github.com/vshigimoto/GoAuthJWTService/internal/services"
	"go.uber.org/zap"
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
