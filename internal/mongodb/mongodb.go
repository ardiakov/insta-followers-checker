package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDbClient struct {
	Client   *mongo.Client
	MongoDsn string
	Database string
}

func InitClient(mongoDsn string, db string) *MongoDbClient {
	ctx := context.Background()

	fmt.Println(mongoDsn)

	credentials := options.Credential{Username: "root", Password: "example"}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDsn).SetAuth(credentials))

	if err != nil {
		panic(err)
	}

	//defer func() {
	//	if err = client.Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()

	mongoDbClient := &MongoDbClient{
		Client:   client,
		MongoDsn: mongoDsn,
		Database: db,
	}

	return mongoDbClient
}

func (r *MongoDbClient) Ping() {
	err := r.Client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		panic(err)
	}
}

func (r *MongoDbClient) Update(collection string, fieldName string, data []string) {
	ctx := context.Background()

	id, _ := primitive.ObjectIDFromHex("630f581c4b2204c4e04f523b")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{fieldName: data}}

	_, err := r.Client.Database("insta").Collection(collection).UpdateOne(ctx, filter, update)

	if err != nil {
		panic(err)
	}

	//fmt.Println(result)
	//
	//result2 := r.Client.Database("insta").Collection("followers").FindOne(ctx, filter)
	//fmt.Println(result2.DecodeBytes())
}

func (r *MongoDbClient) GetDiff(field1 string, field2 string) {
	ctx := context.Background()

	id, _ := primitive.ObjectIDFromHex("630f581c4b2204c4e04f523b")

	pipeline := bson.M{
		"$projectId": bson.M{
			"_id": id,
			"c": bson.M{
				"$setDifference": bson.A{"followers", "followings"},
			},
		},
	}

	//match := bson.M{{"$project", bson.D{{"_id", id}}}}

	//[{ "$project": {"_id": "630f581c4b2204c4e04f523b",  "c": { "$setDifference": [ "$data",[1,2,5] ] } } }]

	cur, err := r.Client.Database("insta").Aggregate(ctx, mongo.Pipeline{pipeline})

	if err != nil {
		panic(err)
	}

	fmt.Println(cur)
}
