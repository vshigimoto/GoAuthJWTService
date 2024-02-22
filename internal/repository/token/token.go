package token

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	db *mongo.Collection
}

func NewRepo(db *mongo.Database, collection string) *Repo {
	return &Repo{
		db: db.Collection(collection),
	}
}
