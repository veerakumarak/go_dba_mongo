package mongo

import (
	"context"
	dba "github.com/veerakumarak/go_dba_core"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func Connect[Entity any](config Config) dba.Repository[Entity] {
	timeout := time.Duration(config.DbTimeOut) * time.Second
	client, err := newMongoClient(config.DbUrl, timeout)
	if err != nil {
		log.Fatal(dba.ErrRepositoryClientNotAbleToConnect)
	}

	return &mongoRepository[Entity]{
		timeout:    timeout,
		collection: client.Database(config.DbName).Collection(config.DbCollection),
	}
}

func newMongoClient(mongoURL string, mongoTimeout time.Duration) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}
