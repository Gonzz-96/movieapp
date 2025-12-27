package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"movieexample.com/gen"
	rating "movieexample.com/rating/internal/controller"
	grpchandler "movieexample.com/rating/internal/handler/grpc"
	"movieexample.com/rating/internal/repository/memory"
)

const serviceName = "rating"

func main() {
	log.Println("Starting the rating service")
	repo := memory.New()
	svc := rating.New(repo)
	h := grpchandler.New(svc)
	lis, err := net.Listen("tcp", "localhost:8082")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	gen.RegisterRatingServiceServer(srv, h)
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}
