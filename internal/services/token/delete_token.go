package token

import (
	"context"
	"fmt"
)

func (s *Service) Delete(ctx context.Context, userGUID string) error {
	err := s.repo.Delete(ctx, userGUID)
	if err != nil {
		return fmt.Errorf("failed to hash refresh token: %v", err)
	}
	return nil
}
