package clients

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var Ctx = context.TODO()

func ConnectMongoDB() error {
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverApi)

	client, err := mongo.Connect(Ctx, opts)
	if err != nil {
		return err
	}

	var result bson.M
	err = client.Database("admin").RunCommand(Ctx, bson.D{{Key: "ping", Value: 1}}).Decode(&result)
	if err != nil {
		panic(err)
	}

	DB = client.Database("todolist")

	return nil
}
