package gRPCFunc

import (
	"OrderService/Service/Configs"
	"OrderService/Service/gRPC_Configs/ProductService"
	"context"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"time"
)

var (
	Search = gRPCProductClient()
)

func gRPCProductClient() ProductService.OrderingServiceClient {
	var opts []grpc.DialOption

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error Load .env File")
	}
	PRODUCTSERVICE := Configs.GetEnv("GRPC_PRODUCTSERVICE")

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(PRODUCTSERVICE, opts...)
	if err != nil {
		log.Fatalln("Fail Dial")
	}
	general := ProductService.NewOrderingServiceClient(conn)
	return general
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

func DecreaseStock(pID string, qty int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_, err := Search.DecreaseStock(ctx, &ProductService.Stock{
		Id:    pID,
		Stock: qty,
	})
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}

func AddSoldProduct(pID string, qty int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_, err := Search.AddSoldProduct(ctx, &ProductService.Stock{
		Id:    pID,
		Stock: qty,
	})
	defer cancel()
	if err != nil {
		return err
	}
	return nil
}

func AddRating(pID string, rating int32, qty int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	_, err := Search.UpdateRating(ctx, &ProductService.Rating{
		Id:          pID,
		TotalRating: rating * qty,
		TotalItem:   qty,
	})
	defer cancel()
	if err != nil {
		return err
	}

	return nil
}
