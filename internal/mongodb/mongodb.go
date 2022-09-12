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

func (r MongoDbClient) CreateDocument() {
	ctx := context.Background()

	doc := bson.D{
		{"followers", ""},
		{"followings", ""},
		{"diff", ""},
	}

	_, err := r.Client.Database("insta").Collection("followers").InsertOne(ctx, doc)

	if err != nil {
		panic(err)
	}
}

func (r *MongoDbClient) GetLastDiff() bson.RawValue {
	ctx := context.Background()

	id, _ := primitive.ObjectIDFromHex("631823a4b1c1589f9c9dc3b3")
	opts := options.FindOne().SetProjection(bson.D{{"diff", 1}})
	result := r.Client.Database(r.Database).Collection("followers").FindOne(ctx, bson.D{{"_id", id}}, opts)

	raw, err := result.DecodeBytes()

	if err != nil {
		panic(err)
	}

	return raw.Lookup("diff")
}

func (r *MongoDbClient) Update(collection string, fieldName string, data []string) {
	ctx := context.Background()

	id, _ := primitive.ObjectIDFromHex("631823a4b1c1589f9c9dc3b3")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{fieldName: data}}

	_, err := r.Client.Database("insta").Collection(collection).UpdateOne(ctx, filter, update)

	if err != nil {
		panic(err)
	}
}

func (r *MongoDbClient) GetDiff(field1 string, field2 string) {
	ctx := context.Background()

	id, _ := primitive.ObjectIDFromHex("631823a4b1c1589f9c9dc3b3")

	match := bson.A{
		bson.M{
			"$match": bson.M{
				"$expr": bson.M{
					"$eq": bson.A{"$_id", id},
				},
			},
		},
		bson.M{
			"$project": bson.M{
				"followings": "$followings",
				"followers":  "$followers",
				"diff": bson.M{
					"$setDifference": bson.A{
						"$followings", "$followers",
					},
				},
			},
		},
		bson.M{
			"$merge": bson.M{
				"into": bson.M{
					"db":   "insta",
					"coll": "followers",
				},
				"on":             "_id",
				"whenMatched":    "replace",
				"whenNotMatched": "insert",
			},
		},
	}

	_, err := r.Client.Database("insta").Collection("followers").Aggregate(ctx, match)

	if err != nil {
		panic(err)
	}
}
