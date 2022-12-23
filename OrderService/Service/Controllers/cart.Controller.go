package Controllers

import (
	"OrderService/Service/Databases"
	"OrderService/Service/Models"
	gRPCFunc "OrderService/Service/gRPC_Configs"
	pb "OrderService/proto"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var DBClient = Databases.DBClient

func AddCart(ctx context.Context, raw *pb.CartItem) error {
	productDetail, err := gRPCFunc.GetProduct(raw.ProductID)
	if err != nil {
		return status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Product Not Found"),
		)
	}

	var cart Models.UserCart
	err = DBClient.Where(&Models.UserCart{
		UserID:    raw.UserID,
		ProductID: raw.ProductID}).First(&cart).Error

	TotalQTY := (cart.QTY) + (raw.Qty)

	if TotalQTY > productDetail.Stock {
		return status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Stock Not Enough"),
		)
	}
	if err != nil {
		tmp := Models.UserCart{}
		if err := DBClient.WithContext(ctx).Create(&Models.UserCart{
			UserID:    raw.UserID,
			ProductID: raw.ProductID,
			QTY:       raw.Qty,
			Total:     productDetail.Price * int64(raw.Qty),
			CreatedAt: time.Now().UnixNano(),
		}).Find(&tmp).Error; err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf(err.Error()),
			)
		}
	} else {
		if err := DBClient.WithContext(ctx).
			Model(&cart).
			Where(&Models.UserCart{
				UserID:    raw.UserID,
				ProductID: raw.ProductID,
			}).
			Updates(Models.UserCart{
				QTY:   cart.QTY + raw.Qty,
				Total: (int64(cart.QTY) + int64(raw.Qty)) * productDetail.Price,
			}).Error; err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf(err.Error()),
			)
		}
	}
	return nil
}

func RemoveCartItem(ctx context.Context, raw *pb.CartItem) error {
	var cart *Models.UserCart

	err := DBClient.WithContext(ctx).Where(&Models.UserCart{
		UserID:    raw.UserID,
		ProductID: raw.ProductID}).First(&cart).Error
	if err != nil {
		return status.Errorf(
			codes.NotFound,
			fmt.Sprintf(err.Error()),
		)
	}
	if err := DBClient.WithContext(ctx).Where(
		&Models.UserCart{
			UserID:    raw.UserID,
			ProductID: raw.ProductID,
		}).
		Delete(&cart).
		Error; err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf(err.Error()),
		)
	}

	return nil
}

func UserCart(ctx context.Context, uID *pb.UserID) (*pb.Cart, error) {
	var cart []*pb.CartItem
	var total int64
	if err := DBClient.WithContext(ctx).Model(&Models.UserCart{}).Where(&Models.UserCart{
		UserID: uID.UserID,
	}).Find(&cart).Select("sum(total)").Scan(&total).Error; err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf(err.Error()),
		)
	}
	return &pb.Cart{
		UserID: uID.UserID,
		Item:   cart,
		Total:  total,
	}, nil
}

func EmptyCart(ctx context.Context, uID *pb.UserID) error {
	var cart *Models.UserCart
	if err := DBClient.WithContext(ctx).Model(&Models.UserCart{}).
		Where(&Models.UserCart{
			UserID: uID.UserID,
		}).
		Delete(&cart).
		Error; err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf(err.Error()),
		)
	}
	return nil
}
