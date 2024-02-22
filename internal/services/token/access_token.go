package token

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/vshigimoto/GoAuthJWTService/internal/entity"
)

const secretKey = "medods"

func (s *Service) GenerateAccessToken(ctx context.Context, userGUID string) (string, error) {
	claims := entity.CustomClaims{
		UserGUID: userGUID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    "auth_service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	accessToken, err := token.SignedString([]byte(secretKey)) // Замените на ваш секретный ключ
	if err != nil {
		return "", fmt.Errorf("failed to sign access token: %v", err)
	}

	return accessToken, nil
}
