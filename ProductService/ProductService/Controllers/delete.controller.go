package Controllers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func DeleteProduct(ctx context.Context, pID string) error {
	product, err := GetProduct(ctx, pID)
	timeNow := time.Now().UnixNano()
	if err != nil {
		return err
	}
	filter := bson.D{bson.E{
		Key:   "_id",
		Value: product.Id,
	}}
	updateData := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "deleted", Value: true},
		bson.E{Key: "updatedAt", Value: timeNow},
	}}}
	if _, err := ProductCollection.UpdateOne(ctx, filter, updateData); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Failed Validate"),
		)
	}
	Eipdate := map[string]interface{}{
		"doc": map[string]interface{}{
			"deleted":   true,
			"updatedAt": time.Now().UnixNano(),
		},
	}
	if err := ElasticUpdate(product.Id, Eipdate, ctx); err != nil {
		return err
	}
	return nil
}

func DecreaseStockProduct(ctx context.Context, pID string, stock int32) error {
	data, err := GetProduct(ctx, pID)
	if err != nil {
		return err
	}
	total := data.Stock - stock
	if total < 0 {
		return status.Errorf(
			codes.Unavailable,
			fmt.Sprintf("Not Enough Stock "),
		)
	}
	filter := bson.D{bson.E{Key: "_id", Value: data.Id}}
	updateData := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "stock", Value: total},
	}}}
	if _, err := ProductCollection.UpdateOne(ctx, filter, updateData); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Failed Validate"),
		)
	}
	Eupdate := map[string]interface{}{
		"doc": map[string]interface{}{
			"stock": total,
		},
	}
	if err := ElasticUpdate(data.Id, Eupdate, ctx); err != nil {
		return err
	}

	return nil
}

func RemoveCategory(ctx context.Context, pID string, arr []string) error {
	_, err := GetProduct(ctx, pID)
	if err != nil {
		return err
	}
	timeNow := time.Now().UnixNano()
	filter := bson.D{bson.E{Key: "_id", Value: pID}}
	updateData := bson.D{
		bson.E{
			Key: "$set", Value: bson.D{
				bson.E{Key: "updatedAt", Value: timeNow},
			},
		},
		bson.E{Key: "$pullAll", Value: bson.M{
			"category": arr,
		}},
	}

	if _, err := ProductCollection.UpdateOne(ctx, filter, updateData); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Failed Validate"),
		)
	}
	data, err := GetProduct(ctx, pID)
	Eupdate := map[string]interface{}{
		"doc": map[string]interface{}{
			"category": data.Category,
			"updateAt": timeNow,
		},
	}
	if err := ElasticUpdate(pID, Eupdate, ctx); err != nil {
		return err
	}
	return nil
}
