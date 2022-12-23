package gRPCFunc

import (
	pb "WebService/App/gRPC_Configs/proto/OrderingService"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

var (
	Cart, Purchase = gRPCOrderingClient()
)

func gRPCOrderingClient() (pb.CartServiceClient, pb.PurchaseClient) {
	var opts []grpc.DialOption

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error Load .env File")
	}
	PRODUCTSERVICE := os.Getenv("GRPC_ORDERINGSERVICE")

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(PRODUCTSERVICE, opts...)
	if err != nil {
		log.Fatalln("Fail Dial")
	}
	cart := pb.NewCartServiceClient(conn)
	purchase := pb.NewPurchaseClient(conn)
	return cart, purchase
}
