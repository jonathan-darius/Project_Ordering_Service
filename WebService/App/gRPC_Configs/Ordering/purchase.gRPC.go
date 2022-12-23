package gRPCFunc

import (
	Ordering "WebService/App/gRPC_Configs/proto/OrderingService"
	"WebService/App/models"
	"context"
	"io"
	"time"
)

func PurchaseCart(uID string) (*Ordering.TransactionID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	res, err := Purchase.PurchaseCart(ctx, &Ordering.UserID{UserID: uID})
	defer cancel()
	if err != nil {
		return nil, err
	}
	return res, err
}

func DetailTransaction(tID string) (*Ordering.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := Purchase.DetailTransaction(ctx, &Ordering.TransactionID{Id: tID})
	defer cancel()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func AddRating(rating *models.ProductRating) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := Purchase.AddRating(ctx, &Ordering.Rating{
		Userid:        rating.UserID,
		TransactionId: rating.TransactionID,
		ProductId:     rating.ProductID,
		Rating:        float64(rating.Rating),
		Desc:          rating.Desc,
	})
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}

func GetAllUserTransaction(uID string, page int32, record int32) ([]*models.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	stream, err := Purchase.PurchaseHistory(ctx, &Ordering.UserPurchase{
		UserID: uID,
		Pagination: &Ordering.Pagination{
			Page:          page,
			RecordPerPage: record,
		},
	})
	defer cancel()
	if err != nil {
		return nil, err
	}
	var data []*models.Transaction
	var detailData []models.TransactionDetail
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		detailData = nil
		for _, transaction := range res.Detail {
			detailData = append(detailData, models.TransactionDetail{
				ProductID: transaction.ProductId,
				QTY:       transaction.Qty,
				Rating:    transaction.Rating,
				Desc:      transaction.Desc,
				Total:     transaction.Total,
			})
		}
		data = append(data, &models.Transaction{
			TransactionID: res.TransactionId,
			UserID:        res.UserId,
			Total:         res.Total,
			CreatedAt:     res.CreateAt,
			Detail:        detailData,
		})
	}
	return data, nil
}
