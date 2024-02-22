package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/vshigimoto/GoAuthJWTService/internal/repository/token"
)

type Repository struct {
	TokenRepo
}

type TokenRepo interface {
	Insert(ctx context.Context, userGuid, hashedToken string) error
	FindOne(ctx context.Context, key string) (token.RefreshTokenDocument, error)
}

func New(db *mongo.Database, collection string) *Repository {
	return &Repository{
		TokenRepo: token.NewRepo(db, collection),
	}
}
