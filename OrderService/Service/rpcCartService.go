package main

import (
	"OrderService/Service/Controllers"
	pb "OrderService/proto"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (*Server) AddCartItem(ctx context.Context, in *pb.CartItem) (*empty.Empty, error) {
	log.Printf("Add Cart Invoked for UserID : %s With ProductID ; %s\n", in.UserID, in.ProductID)
	err := Controllers.AddCart(ctx, in)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (*Server) RemoveCartItem(ctx context.Context, in *pb.CartItem) (*empty.Empty, error) {
	log.Printf("Remove Cart Invoked for UserID : %s With ProductID ; %s\n", in.UserID, in.ProductID)
	err := Controllers.RemoveCartItem(ctx, in)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (*Server) GetCart(ctx context.Context, in *pb.UserID) (*pb.Cart, error) {
	log.Printf("Get Cart Invoked for UserID : %s \n", in.UserID)
	res, err := Controllers.UserCart(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (*Server) EmptyCart(ctx context.Context, in *pb.UserID) (*empty.Empty, error) {
	log.Printf("Empty Cart Invoked for UserID : %s \n", in.UserID)
	err := Controllers.EmptyCart(ctx, in)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
