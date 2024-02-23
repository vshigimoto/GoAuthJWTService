package token

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) HashToken(ctx context.Context, refreshToken string) (string, error) {
	_ = ctx
	if len(refreshToken) == 0 {
		return "", fmt.Errorf("token must not be empty")
	}
	hashedRefreshToken, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash refresh token: %v", err)
	}
	return string(hashedRefreshToken), nil
}
