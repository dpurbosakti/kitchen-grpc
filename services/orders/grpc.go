package main

import (
	"log"
	"net"

	handler "github.com/dpurbosakti/kitchen-grpc/services/orders/handler/orders"
	"github.com/dpurbosakti/kitchen-grpc/services/orders/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{
		addr: addr,
	}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)
	// register grpc services
	orderService := service.NewOrderService()
	handler.NewGRPCOrdersService(grpcServer, orderService)

	log.Println("Starting gRPC server on", s.addr)

	return grpcServer.Serve(lis)
}
