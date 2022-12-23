package Databases

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var (
	Client *mongo.Client = DBInstance()
)

func DBInstance() *mongo.Client {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error Load .env File")
	}
	MongoDb := os.Getenv("MONGODB_URL")
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
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

func OpenCollection(client *mongo.Client, colName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("Mini_Product_Ordering").Collection(colName)
	return collection
}
