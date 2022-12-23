package Controllers

import (
	"ProductService/ProductService/Utils"
	pb "ProductService/proto"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func UpdateProduct(ctx context.Context, In *pb.UpdateProductMSG) error {
	_, err := GetProduct(ctx, In.Id)
	if err != nil {
		return err
	}
	var query bson.D
	timeNow := time.Now().UnixNano()
	tmp := map[string]interface{}{
		"updatedAt": timeNow,
	}

	filter := bson.D{bson.E{Key: "_id", Value: In.Id}}

	query = append(query, bson.E{Key: "updatedAt", Value: timeNow})
	if In.Name.Check {
		query = append(query, bson.E{Key: "name", Value: In.Name})
		tmp["name"] = In.Name
	}
	if In.Price.Check {
		query = append(query, bson.E{Key: "price", Value: In.Price.Value})
		tmp["price"] = In.Price.Value
	}
	if In.Stock.Check {
		query = append(query, bson.E{Key: "stock", Value: In.Stock.Value})
		tmp["stock"] = In.Stock.Value
	}
	if In.Category.Check {
		setCategory := Utils.Set(In.Category.Value)
		query = append(query, bson.E{Key: "category", Value: setCategory.Value()})
		tmp["category"] = setCategory.Value()
	}
	if In.Image.Check {
		query = append(query, bson.E{Key: "image", Value: In.Image.Value})
		tmp["image"] = In.Image.Value
	}
	Eupdate := map[string]interface{}{
		"doc": tmp,
	}

	updateData := bson.D{bson.E{Key: "$set", Value: query}}
	if _, err := ProductCollection.UpdateOne(ctx, filter, updateData); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Failed Validate"),
		)
	}

	if err := ElasticUpdate(In.Id, Eupdate, ctx); err != nil {
		return err
	}
	return nil

}
