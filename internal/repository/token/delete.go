package token

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repo) Delete(ctx context.Context, key string) error {
	filter := bson.M{"user_guid": key}

	_, err := r.db.DeleteOne(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("no document found %s : %v", key, key)
		}
		return fmt.Errorf("failure retrieving from mongoDB: %v", err)
	}

	return nil
}
