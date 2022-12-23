package Controllers

import (
	"OrderService/Service/Models"
	gRPCFunc "OrderService/Service/gRPC_Configs"
	pb "OrderService/proto"
	"context"
	"fmt"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func PurchaseCart(ctx context.Context, uID *pb.UserID) (*pb.TransactionID, error) {
	newID := uuid.New().String()
	cart, err := UserCart(ctx, uID)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf(err.Error()),
		)
	}

	dataTransaction := &Models.Purchase{
		TransactionID: newID,
		UserID:        uID.UserID,
		Total:         cart.Total,
		CreatedAt:     time.Now().UnixNano(),
	}

	var product []*Models.PurchaseDetail
	for _, item := range cart.Item {
		product = append(product, &Models.PurchaseDetail{
			TransactionID: newID,
			ProductID:     item.ProductID,
			QTY:           item.Qty,
			Total:         item.Total,
			Rating:        -1,
		})
		productDetail, err := gRPCFunc.GetProduct(item.ProductID)
		if err != nil {
			return nil, status.Errorf(
				codes.NotFound,
				fmt.Sprintf("Product Not Found"),
			)
		}
		if productDetail.Stock < item.Qty {
			return nil, status.Errorf(
				codes.NotFound,
				fmt.Sprintf("Not Enough Stock"),
			)
		}
	}
	for _, item := range product {
		if err := gRPCFunc.DecreaseStock(item.ProductID, item.QTY); err != nil {
			return nil, err
		}
		if err := gRPCFunc.AddSoldProduct(item.ProductID, item.QTY); err != nil {
			return nil, err
		}
	}

	err = DBClient.Model(&Models.Purchase{}).Create(dataTransaction).Error
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf(err.Error()),
		)
	}

	err = DBClient.Model(&Models.PurchaseDetail{}).Create(&product).Error
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf(err.Error()),
		)
	}

	if err := EmptyCart(ctx, uID); err != nil {
		return nil, err
	}

	return &pb.TransactionID{
		Id: newID,
	}, nil
}

func GetPurchaseHistory(uID string, page int64, record int64) ([]*pb.Transaction, error) {
	var transaction []*Models.Purchase
	offset := (page - 1) * record
	query := DBClient.WithContext(context.Background()).Limit(int(record)).Offset(int(offset))
	if err := query.Model(&Models.Purchase{}).Where(&Models.Purchase{
		UserID: uID,
	}).Find(&transaction).Error; err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf(err.Error()),
		)
	}

	var clean []*pb.Transaction

	for _, purchase := range transaction {
		transaction, err := PurchaseData(purchase.TransactionID)
		if err != nil {
			return nil, err
		}

		detail, err := PurchaseDetail(purchase.TransactionID)
		if err != nil {
			return nil, err
		}
		clean = append(clean, &pb.Transaction{
			TransactionId: transaction.TransactionID,
			UserId:        transaction.UserID,
			Total:         transaction.Total,
			CreateAt:      transaction.CreatedAt,
			Detail:        detail,
		})
	}

	return clean, nil
}

func PurchaseData(tID string) (*Models.Purchase, error) {
	var data *Models.Purchase
	if err := DBClient.WithContext(context.Background()).Model(Models.Purchase{}).Where(&Models.Purchase{
		TransactionID: tID,
	}).Find(&data).Error; err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf(err.Error()),
		)
	}
	return data, nil
}

func PurchaseDetail(tID string) ([]*pb.DetailTransaction, error) {
	var data []*Models.PurchaseDetail
	if err := DBClient.WithContext(context.Background()).Model(Models.PurchaseDetail{}).Where(&Models.PurchaseDetail{
		TransactionID: tID,
	}).Find(&data).Error; err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf(err.Error()),
		)
	}

	var final []*pb.DetailTransaction
	for _, product := range data {
		final = append(final, &pb.DetailTransaction{
			ProductId: product.ProductID,
			Qty:       product.QTY,
			Rating:    product.Rating,
			Desc:      product.Desc,
			Total:     product.Total,
		})
	}
	return final, nil
}

func AddRatingProduct(ctx context.Context, in *pb.Rating) error {
	transaction, err := PurchaseData(in.TransactionId)
	if err != nil {
		return err
	}
	if transaction.UserID != in.Userid {
		return status.Errorf(
			codes.PermissionDenied,
			fmt.Sprintf(err.Error()),
		)
	}
	var newData *Models.PurchaseDetail
	if err := DBClient.WithContext(ctx).Model(&Models.PurchaseDetail{}).Where(&Models.PurchaseDetail{
		TransactionID: in.TransactionId,
		ProductID:     in.ProductId,
		Rating:        -1,
	}).Find(&newData).Error; err != nil {
		return status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf(err.Error()),
		)
	}

	if err := DBClient.WithContext(ctx).Model(&Models.PurchaseDetail{}).Where(&Models.PurchaseDetail{
		TransactionID: in.TransactionId,
		ProductID:     in.ProductId,
		Rating:        -1,
	}).Updates(&Models.PurchaseDetail{
		Rating: in.Rating,
		Desc:   in.Desc,
	}).Find(&newData).Error; err != nil {
		return status.Errorf(
			codes.AlreadyExists,
			fmt.Sprintf(err.Error()),
		)
	}
	if err = gRPCFunc.AddRating(in.ProductId, int32(in.Rating), newData.QTY); err != nil {
		return err
	}
	return nil
}
