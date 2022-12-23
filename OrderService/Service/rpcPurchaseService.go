package main

import (
	"OrderService/Service/Controllers"
	pb "OrderService/proto"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (*Server) PurchaseCart(ctx context.Context, in *pb.UserID) (*pb.TransactionID, error) {
	log.Printf("Purchase Cart Invoked for UserID : %s \n", in.UserID)
	res, err := Controllers.PurchaseCart(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (*Server) PurchaseHistory(in *pb.UserPurchase, stream pb.Purchase_PurchaseHistoryServer) error {
	log.Printf("Purchase History Invoked for UserID : %s \n", in.UserID)
	recordPerPage := in.Pagination.RecordPerPage
	page := in.Pagination.Page

	if recordPerPage == 0 {
		recordPerPage = 5
	}
	if page == 0 {
		page = 1
	}

	data, err := Controllers.GetPurchaseHistory(in.UserID, int64(page), int64(recordPerPage))
	if err != nil {
		return err
	}
	for _, transaction := range data {
		err := stream.Send(&pb.Transaction{
			TransactionId: transaction.TransactionId,
			UserId:        transaction.UserId,
			Total:         transaction.Total,
			CreateAt:      transaction.CreateAt,
			Detail:        transaction.Detail,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (*Server) DetailTransaction(_ context.Context, in *pb.TransactionID) (*pb.Transaction, error) {
	log.Printf("Purchase Detail Invoked for Transaction : %s \n", in.Id)

	transaction, err := Controllers.PurchaseData(in.Id)
	if err != nil {
		return nil, err
	}

	detail, err := Controllers.PurchaseDetail(in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Transaction{
		TransactionId: in.Id,
		UserId:        transaction.UserID,
		Total:         transaction.Total,
		CreateAt:      transaction.CreatedAt,
		Detail:        detail,
	}, nil
}

func (*Server) AddRating(ctx context.Context, in *pb.Rating) (*empty.Empty, error) {
	log.Printf("Add Rating Invoked for Transaction : %s \n", in.TransactionId)
	err := Controllers.AddRatingProduct(ctx, in)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
