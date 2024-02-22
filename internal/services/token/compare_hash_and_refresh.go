package token

import (
	"context"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CompareHashAndRefresh(ctx context.Context, refreshToken string) (string, error) {

	HashedToken, err := s.HashToken(ctx, refreshToken)

	hash, err := s.repo.FindOne(ctx, HashedToken)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash.HashedToken), []byte(refreshToken)); err != nil {
		return "", err
	}
	return hash.UserGuid, nil
}
