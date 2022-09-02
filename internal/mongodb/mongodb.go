package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDbClient struct {
	Client   *mongo.Client
	MongoDsn string
	Database string
}

func InitClient(mongoDsn string, db string) *MongoDbClient {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDsn))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	mongoDbClient := &MongoDbClient{
		Client:   client,
		MongoDsn: mongoDsn,
		Database: db,
	}

	return mongoDbClient
}

func (r *MongoDbClient) Ping() {
}

func (r *MongoDbClient) Insert() {
	r.Client.Database(r.Database).Collection()
}
