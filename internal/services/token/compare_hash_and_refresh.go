package token

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CompareHashAndRefresh(ctx context.Context, refreshToken, userGUID string) error {
	hash, err := s.repo.FindOne(ctx, userGUID)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash.HashedToken), []byte(refreshToken)); err != nil {
		return err
	}
	return nil
}
