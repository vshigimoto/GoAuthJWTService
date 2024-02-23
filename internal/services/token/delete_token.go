package token

import (
	"context"
	"fmt"
)

func (s *Service) Delete(ctx context.Context, userGuid string) error {
	err := s.repo.Delete(ctx, userGuid)
	if err != nil {
		return fmt.Errorf("failed to hash refresh token: %v", err)
	}
	return nil
}
