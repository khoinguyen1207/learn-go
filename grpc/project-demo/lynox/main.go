package main

import (
	"context"
	"log"
	"lynox/handler"
	"lynox/repository"
	"lynox/service"
	"net"
	"net/http"

	pb "proto-demo-repo/gen/lynox"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	grpcPort := ":50051"
	httpPort := ":8080"

	// --- gRPC server ---
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewUserRepository()
	srv := service.NewUserService(repo)
	handler := handler.NewUserHandler(srv)
	grpcServer := grpc.NewServer()
	pb.RegisterLynoxServer(grpcServer, handler)

	go func() {
		log.Println("gRPC running on", grpcPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	// --- HTTP Gateway ---
	ctx := context.Background()
	mux := runtime.NewServeMux()

	err = pb.RegisterLynoxHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:50051",
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP running on", httpPort)
	http.ListenAndServe(httpPort, mux)
}
