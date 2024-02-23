package repository

import (
	"context"

	"github.com/vshigimoto/GoAuthJWTService/internal/repository/token"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	TokenRepo
}

type TokenRepo interface {
	Insert(ctx context.Context, userGUID, hashedToken string) error
	FindOne(ctx context.Context, key string) (token.RefreshTokenDocument, error)
	Delete(ctx context.Context, key string) error
}

func New(db *mongo.Database, collection string) *Repository {
	return &Repository{
		TokenRepo: token.NewRepo(db, collection),
	}
}
