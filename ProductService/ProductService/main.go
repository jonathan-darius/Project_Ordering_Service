package main

// cd protoc && protoc --go_out=../proto/. --go_opt=paths=source_relative --go-grpc_out=../proto/. --go-grpc_opt=paths=source_relative product.proto && cd ..
import (
	"ProductService/ProductService/Config"
	pb "ProductService/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Server struct {
	pb.ProductManagementServiceServer
	pb.OrderingServiceServer
}

var (
	addr       = Config.GetEnv("SERVER_ADDR")
	lis        net.Listener
	err        error
	gRPCServer *grpc.Server
)

func init() {
	var opts []grpc.ServerOption
	lis, err = net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	gRPCServer = grpc.NewServer(opts...)
	log.Printf("Listening at %s\n", addr)
	//Databases.OpenCollection("user")
}

// main ToDo SetUp SSL & Dockerfile
func main() {
	pb.RegisterProductManagementServiceServer(gRPCServer, &Server{})
	pb.RegisterOrderingServiceServer(gRPCServer, &Server{})

	reflection.Register(gRPCServer)
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
