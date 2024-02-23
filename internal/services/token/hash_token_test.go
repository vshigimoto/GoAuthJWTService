package token

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestHashToken(t *testing.T) {
	s := Service{}

	testToken := "mySecretRefreshToken"
	hashedToken, err := s.HashToken(context.Background(), testToken)
	if err != nil {
		t.Fatalf("HashToken() error = %v, wantErr = false", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedToken), []byte(testToken))
	if err != nil {
		t.Errorf("Hash does not match the original token. err = %v", err)
	}
}

func TestHashTokenEmpty(t *testing.T) {
	s := &Service{}

	_, err := s.HashToken(context.Background(), "")
	if err == nil {
		t.Errorf("HashToken() with empty token did not fail")
	}
}
