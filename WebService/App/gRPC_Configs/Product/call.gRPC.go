package gRPCFunc

import (
	"WebService/App/gRPC_Configs/proto/ProductService"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

var (
	Management, Search = gRPCProductClient()
)

func gRPCProductClient() (ProductService.ProductManagementServiceClient, ProductService.OrderingServiceClient) {
	var opts []grpc.DialOption

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error Load .env File")
	}
	PRODUCTSERVICE := os.Getenv("GRPC_PRODUCTSERVICE")

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(PRODUCTSERVICE, opts...)
	if err != nil {
		log.Fatalln("Fail Dial")
	}
	management := ProductService.NewProductManagementServiceClient(conn)
	general := ProductService.NewOrderingServiceClient(conn)
	return management, general
}
