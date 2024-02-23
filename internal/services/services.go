package services

import (
	"context"

	"github.com/vshigimoto/GoAuthJWTService/internal/repository"
	"github.com/vshigimoto/GoAuthJWTService/internal/services/token"
)

type Services struct {
	TokenService
}

type TokenService interface {
	GenerateAccessToken(ctx context.Context, userGUID string) (string, error)
	GenerateRefreshToken(ctx context.Context, userGUID string) (string, error)
	CompareHashAndRefresh(ctx context.Context, refreshToken, userGUID string) error
	HashToken(ctx context.Context, refreshToken string) (string, error)
	Delete(ctx context.Context, userGUID string) error
}

func New(repo *repository.Repository) *Services {
	return &Services{
		TokenService: token.NewService(repo.TokenRepo),
	}
}
