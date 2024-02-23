package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	duration = time.Minute
)

func Init(url, username, password string) (*mongo.Client, error) {
	credentials := options.Credential{
		Username: username,
		Password: password,
	}

	clientOptions := options.Client().ApplyURI(url).SetAuth(credentials)

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	return client, err
}
