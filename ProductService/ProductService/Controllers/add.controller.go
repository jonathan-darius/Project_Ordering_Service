package Controllers

import (
	"ProductService/ProductService/Models"
	"ProductService/ProductService/Utils"
	pb "ProductService/proto"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"time"
)

func AddProduct(ctx context.Context, in *pb.Product) (*Models.ProductModel, error) {
	timeNow := time.Now().UnixNano()
	data := Models.ProductModel{
		Id:        primitive.NewObjectID().Hex(),
		Name:      in.Name,
		Stock:     in.Stock,
		Price:     in.Price,
		Sold:      0,
		Rating:    0,
		Rated:     0,
		Desc:      in.Desc,
		Category:  in.Category,
		Image:     []string{},
		CreatedBy: in.CreatedBy,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
		Deleted:   false,
	}
	if in.Category == nil {
		data.Category = []string{}
	}

	if _, err := ProductCollection.InsertOne(ctx, data); err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	Edata, err := json.Marshal(data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	res, err := EClient.Index(
		"mini-product-ordering",               // Index name
		strings.NewReader(string(Edata)),      // Document body
		EClient.Index.WithDocumentID(data.Id), // Document ID
		EClient.Index.WithRefresh("true"),     // Refresh
	)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	defer res.Body.Close()
	return &data, nil
}

// AddStockProduct ToDo Update elastic stock
func AddStockProduct(ctx context.Context, pID string, stock int32) error {
	data, err := GetProduct(ctx, pID)
	total := data.Stock + stock
	timeNow := time.Now().UnixNano()
	if err != nil {
		return status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Product Not Found"),
		)
	}
	filter := bson.D{bson.E{
		Key:   "_id",
		Value: data.Id,
	}}
	updateData := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "updatedAt", Value: timeNow},
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
			"stock":     total,
			"updatedAt": timeNow,
		},
	}
	if err := ElasticUpdate(data.Id, Eupdate, ctx); err != nil {
		return err
	}

	return nil
}

func AddSold(ctx context.Context, pID string, stock int32) error {
	data, err := GetProduct(ctx, pID)
	if err != nil {
		return err
	}
	total := data.Sold + stock
	filter := bson.D{bson.E{Key: "_id", Value: data.Id}}
	updateData := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "sold", Value: total},
	}}}

	if _, err := ProductCollection.UpdateOne(ctx, filter, updateData); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Failed Validate"),
		)
	}
	Eupdate := map[string]interface{}{
		"doc": map[string]interface{}{
			"sold": total,
		},
	}
	if err := ElasticUpdate(data.Id, Eupdate, ctx); err != nil {
		return err
	}

	return nil
}

func AddRating(ctx context.Context, pID string, TotalRating int32, TotalItem int32) error {
	data, err := GetProduct(ctx, pID)
	if err != nil {
		return err
	}

	oldRatingTotal := data.Rating * float64(data.Rated)
	newtotalRating := data.Rated + TotalItem

	newRating := Utils.RoundFloat((oldRatingTotal+float64(TotalRating))/float64(newtotalRating), 2)
	filter := bson.D{bson.E{Key: "_id", Value: data.Id}}
	updateData := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "rating", Value: newRating},
		bson.E{Key: "rated", Value: newtotalRating},
	}}}

	if _, err := ProductCollection.UpdateOne(ctx, filter, updateData); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Failed Validate"),
		)
	}
	Eupdate := map[string]interface{}{
		"doc": map[string]interface{}{
			"rating": newRating,
			"rated":  newtotalRating,
		},
	}
	if err := ElasticUpdate(data.Id, Eupdate, ctx); err != nil {
		return err
	}
	return nil
}

func AddCategory(ctx context.Context, pID string, arr []string) error {
	data, err := GetProduct(ctx, pID)
	if err != nil {
		return err
	}

	timeNow := time.Now().UnixNano()

	newCategory := Utils.Set(append(arr, data.Category...)).Value()

	filter := bson.D{bson.E{Key: "_id", Value: pID}}
	updateData := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "category", Value: newCategory},
		bson.E{Key: "updatedAt", Value: timeNow},
	}}}
	if _, err := ProductCollection.UpdateOne(ctx, filter, updateData); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Failed Validate"),
		)
	}
	Eupdate := map[string]interface{}{
		"doc": map[string]interface{}{
			"category": newCategory,
			"updateAt": timeNow,
		},
	}
	if err := ElasticUpdate(pID, Eupdate, ctx); err != nil {
		return err
	}
	return nil
}

func AddImage(ctx context.Context, pID string, imgInfo *Utils.ImageInfo) error {
	timeNow := time.Now().UnixNano()
	filter := bson.D{bson.E{
		Key:   "_id",
		Value: pID,
	}}

	updateData := bson.D{
		bson.E{Key: "$set", Value: bson.D{
			bson.E{Key: "updatedAt", Value: timeNow},
		}},
		bson.E{Key: "$push", Value: bson.M{
			"image": imgInfo.Path,
		}},
	}
	if _, err := ProductCollection.UpdateOne(ctx, filter, updateData); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Failed Update"),
		)
	}
	data, err := GetProduct(ctx, pID)
	if err != nil {
		return err
	}

	Eupdate := map[string]interface{}{
		"doc": map[string]interface{}{
			"image":    data.Image,
			"updateAt": timeNow,
		},
	}
	if err := ElasticUpdate(pID, Eupdate, ctx); err != nil {
		return err
	}

	return nil
}
