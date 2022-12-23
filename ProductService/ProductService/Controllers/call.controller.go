package Controllers

import (
	"ProductService/ProductService/Models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func GetProduct(ctx context.Context, pID string) (*Models.ProductModel, error) {
	var product *Models.ProductModel
	res := ProductCollection.FindOne(ctx, bson.M{"_id": pID, "deleted": false})
	if err := res.Decode(&product); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Product Not Found"),
		)
	}
	return product, nil
}

func GetAllProduct(record int32, page int32) (*mongo.Cursor, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	l := int64(record)
	skip := int64(page*record - record)
	fOpt := options.FindOptions{Limit: &l, Skip: &skip}
	data, err := ProductCollection.Find(ctx, bson.D{
		{"deleted", false},
	}, &fOpt)
	defer cancel()
	if err != nil {
		return nil, err
	}
	return data, nil
}
