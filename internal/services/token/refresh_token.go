package token

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func (s *Service) GenerateRefreshToken(ctx context.Context, userGuid string) (string, error) {
	refreshTokenBytes := make([]byte, 32)
	_, err := rand.Read(refreshTokenBytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate refresh token: %v", err)
	}
	refreshToken := base64.StdEncoding.EncodeToString(refreshTokenBytes)

	hashedRefreshToken, err := s.HashToken(ctx, refreshToken)
	if err != nil {
		return "", fmt.Errorf("failed to generate hash refreshToken: %v", err)
	}

	err = s.repo.Insert(ctx, userGuid, hashedRefreshToken)
	if err != nil {
		return "", fmt.Errorf("failed to insert refresh token: %v", err)
	}

	return refreshToken, nil
}
