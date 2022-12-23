package gRPCFunc

import (
	Ordering "WebService/App/gRPC_Configs/proto/OrderingService"
	"WebService/App/models"
	"context"
	"time"
)

func AddCart(data models.CartItemSend) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_, err := Cart.AddCartItem(ctx, &Ordering.CartItem{
		UserID:    data.UserID,
		ProductID: data.ProductID,
		Qty:       data.QTY,
		Total:     data.Total,
	})
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}

func RemoveCartItem(data models.CartItemSend) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_, err := Cart.RemoveCartItem(ctx, &Ordering.CartItem{
		UserID:    data.UserID,
		ProductID: data.ProductID,
		Qty:       data.QTY,
		Total:     data.Total,
	})
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}

func GetUserCart(uID string) (*Ordering.Cart, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	res, err := Cart.GetCart(ctx, &Ordering.UserID{UserID: uID})
	defer cancel()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func EmptyUserCart(uID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_, err := Cart.EmptyCart(ctx, &Ordering.UserID{UserID: uID})
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}
