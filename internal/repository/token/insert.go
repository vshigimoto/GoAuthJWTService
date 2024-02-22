package token

import (
	"context"
	"fmt"
	"time"
)

type RefreshTokenDocument struct {
	HashedToken  string    `bson:"hashed_token"`
	UserGuid     string    `bson:"user_guid"`
	CreationDate time.Time `bson:"creation_date"`
}

func (r *Repo) Insert(ctx context.Context, userGuid, hashedToken string) error {

	doc := RefreshTokenDocument{
		HashedToken:  hashedToken,
		UserGuid:     userGuid,
		CreationDate: time.Now(),
	}

	_, err := r.db.InsertOne(ctx, doc)
	if err != nil {
		return fmt.Errorf("failure insert to mongoDB: %v", err)
	}

	return nil
}
