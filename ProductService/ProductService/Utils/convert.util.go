package Utils

import (
	"ProductService/ProductService/Models"
	pb "ProductService/proto"
)

func ProductDocToProduct(data *Models.ProductModel) *pb.Product {
	return &pb.Product{
		Id:        data.Id,
		Name:      data.Name,
		Price:     data.Price,
		Stock:     data.Stock,
		Rating:    data.Rating,
		Rated:     data.Rated,
		Sold:      data.Sold,
		Category:  data.Category,
		Desc:      data.Desc,
		Image:     data.Image,
		CreatedBy: data.CreatedBy,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		Deleted:   data.Deleted,
	}
}
