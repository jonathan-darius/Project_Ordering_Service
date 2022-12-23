package gRPCFunc

import (
	"WebService/App/gRPC_Configs/proto/ProductService"
	"WebService/App/models"
	"context"
	"io"
	"time"
)

func SearchProduct(data *models.SearchProduct, page int32, record int32) ([]*models.ProductShow, error) {
	query := ProductService.SearchQuery{
		Keyword:   data.Keyword,
		Category:  data.Category,
		PriceLow:  data.PriceLow,
		PriceHigh: data.PriceHigh,
		Rating:    data.Rating,
		SortBy:    data.SortBy,
		Order:     data.Order,
		Pagination: &ProductService.Pagination{
			Page:          page,
			RecordPerPage: record,
		},
	}
	ctx := context.Background()
	stream, err := Search.SearchProduct(ctx, &query)
	if err != nil {
		return nil, err
	}
	var product []*models.ProductShow
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		product = append(product, &models.ProductShow{
			ID:        res.Id,
			Name:      res.Name,
			Price:     res.Price,
			Stock:     res.Stock,
			Sold:      res.Sold,
			Rating:    res.Rating,
			Rated:     res.Rated,
			Desc:      res.Desc,
			Category:  res.Category,
			CreatedAt: res.CreatedAt,
		})
	}
	return product, nil
}

func GetProduct(pID string) (*ProductService.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	res, err := Search.GetProductByID(ctx, &ProductService.ProductId{Id: pID})
	defer cancel()
	if err != nil {
		return nil, err
	}
	return res, nil
}
