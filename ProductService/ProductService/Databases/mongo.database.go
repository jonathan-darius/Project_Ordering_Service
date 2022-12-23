package Databases

import (
	"ProductService/ProductService/Config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	Client    *mongo.Client = DBInstance()
	MONGO_URI               = Config.GetEnv("MONGODB_URL")
)

func DBInstance() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo Connected")
	return client
}

func OpenCollection(colName string) *mongo.Collection {
	var collection *mongo.Collection = Client.Database("Mini_Product_Ordering").Collection(colName)
	return collection
}
