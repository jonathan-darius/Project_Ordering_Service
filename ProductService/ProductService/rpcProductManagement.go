package main

import (
	"ProductService/ProductService/Controllers"
	pb "ProductService/proto"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (*Server) RegisterProduct(ctx context.Context, in *pb.Product) (*pb.ProductId, error) {
	log.Println("Product Register Invoke")
	res, err := Controllers.AddProduct(ctx, in)
	if err != nil {
		return nil, err
	}
	return &pb.ProductId{Id: res.Id}, err
}

func (*Server) DeleteProduct(ctx context.Context, in *pb.ProductId) (*empty.Empty, error) {
	log.Printf("Delete Product was invoked with %v\n", in)
	err := Controllers.DeleteProduct(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (*Server) AddStock(ctx context.Context, in *pb.Stock) (*empty.Empty, error) {
	log.Printf("Add Product Stock was invoked with %v\n", in)
	err := Controllers.AddStockProduct(ctx, in.Id, in.Stock)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (*Server) UpdateProduct(ctx context.Context, in *pb.UpdateProductMSG) (*empty.Empty, error) {
	log.Println("Product Update Invoke")
	if err := Controllers.UpdateProduct(ctx, in); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (*Server) AddCategory(ctx context.Context, in *pb.UpdateProductMSG) (*empty.Empty, error) {
	log.Println("Add Category Invoke")
	if err := Controllers.AddCategory(ctx, in.Id, in.Category.Value); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (*Server) RemoveCatogory(ctx context.Context, in *pb.UpdateProductMSG) (*empty.Empty, error) {
	log.Println("Remove Category Invoke")
	if err := Controllers.RemoveCategory(ctx, in.Id, in.Category.Value); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
