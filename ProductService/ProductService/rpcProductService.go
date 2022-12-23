package main

import (
	"ProductService/ProductService/Controllers"
	"ProductService/ProductService/Models"
	"ProductService/ProductService/Utils"
	pb "ProductService/proto"
	"context"
	"encoding/json"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (*Server) GetProductByID(ctx context.Context, in *pb.ProductId) (*pb.Product, error) {
	log.Printf("Get Product By ID was invoked with %v\n", in)
	data, err := Controllers.GetProduct(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return Utils.ProductDocToProduct(data), nil
}

func (*Server) GetAllProduct(in *pb.Pagination, stream pb.OrderingService_GetAllProductServer) error {
	log.Println("Get All Product was invoked")
	recordPerPage := in.RecordPerPage
	page := in.Page
	data, err := Controllers.GetAllProduct(recordPerPage, page)
	if err != nil {
		return err
	}
	var product Models.ProductModel
	for data.Next(context.Background()) {
		if err := data.Decode(&product); err != nil {
			return err
		}
		err := stream.Send(Utils.ProductDocToProduct(&product))
		if err != nil {
			return err
		}
	}

	return nil
}

func (*Server) DecreaseStock(ctx context.Context, in *pb.Stock) (*empty.Empty, error) {
	log.Printf("Decrease Product Stock was invoked with %v\n", in)
	err := Controllers.DecreaseStockProduct(ctx, in.Id, in.Stock)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (*Server) AddSoldProduct(ctx context.Context, in *pb.Stock) (*empty.Empty, error) {
	log.Printf("Add Sold Product was invoked with %v\n", in)
	err := Controllers.AddSold(ctx, in.Id, in.Stock)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (*Server) UpdateRating(ctx context.Context, in *pb.Rating) (*empty.Empty, error) {
	log.Printf("Add Rating Product was invoked with %v\n", in)
	err := Controllers.AddRating(ctx, in.Id, in.TotalRating, in.TotalItem)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (*Server) SearchProduct(in *pb.SearchQuery, stream pb.OrderingService_SearchProductServer) error {
	log.Println("Search Product was invoked")
	recordPerPage := in.Pagination.RecordPerPage
	page := in.Pagination.Page
	if recordPerPage <= 0 {
		recordPerPage = 5
	}
	if page <= 1 {
		page = 1
	}
	hitsData, err := Controllers.Search(recordPerPage, page, in)
	if err != nil {
		return err
	}
	var product Models.ProductModel
	for _, hit := range hitsData["hits"].(map[string]interface{})["hits"].([]interface{}) {
		raw, err := json.Marshal(hit.(map[string]interface{})["_source"].(map[string]interface{}))
		if err != nil {
			return err
		}
		if err := json.Unmarshal(raw, &product); err != nil {
			return err
		}
		if err = stream.Send(Utils.ProductDocToProduct(&product)); err != nil {
			return err
		}
	}
	return nil
}

func (*Server) UploadProductImage(stream pb.ProductManagementService_UploadProductImageServer) error {
	err := Controllers.StreamImage(stream)
	if err != nil {
		return err
	}

	return nil
}
