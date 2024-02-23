package token

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repo) FindOne(ctx context.Context, key string) (RefreshTokenDocument, error) {
	var result RefreshTokenDocument

	filter := bson.M{"user_guid": key}

	err := r.db.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return RefreshTokenDocument{}, fmt.Errorf("no document found %s : %v", key, key)
		}
		return RefreshTokenDocument{}, fmt.Errorf("failure retrieving from mongoDB: %v", err)
	}

	return result, nil
}
