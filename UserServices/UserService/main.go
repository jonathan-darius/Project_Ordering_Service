package main

// protoc --go_out=../proto/. --go_opt=paths=source_relative --go-grpc_out=../proto/. --go-grpc_opt=paths=source_relative user.proto
import (
	pb "UserServices/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var addr = "0.0.0.0:50051"

// main ToDo SetUp SSL & Dockerfile
func main() {
	var opts []grpc.ServerOption

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening at %s\n", addr)

	gRPCServer := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(gRPCServer, &Server{})
	//pb.RegisterAuthServiceServer(gRPCServer, &Server{})
	reflection.Register(gRPCServer)

	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
