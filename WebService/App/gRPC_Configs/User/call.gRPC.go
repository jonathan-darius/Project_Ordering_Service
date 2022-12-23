package gRPCFunc

import (
	"WebService/App/gRPC_Configs/proto/UserService"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

var (
	Client UserService.UserServiceClient = gRPCUserClient()
)

func gRPCUserClient() UserService.UserServiceClient {
	var opts []grpc.DialOption

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error Load .env File")
	}
	USERSERVICE := os.Getenv("GRPC_USERSERVICE")

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(USERSERVICE, opts...)
	if err != nil {
		log.Fatalln("Fail Dial")
	}
	client := UserService.NewUserServiceClient(conn)
	return client
}
