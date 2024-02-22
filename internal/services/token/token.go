package token

import "github.com/vshigimoto/GoAuthJWTService/internal/repository"

type Service struct {
	repo repository.TokenRepo
}

func NewService(repo repository.TokenRepo) *Service {
	return &Service{
		repo: repo,
	}
}
