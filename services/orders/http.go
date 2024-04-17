package main

import (
	"log"
	"net/http"

	handler "github.com/dpurbosakti/kitchen-grpc/services/orders/handler/orders"
	"github.com/dpurbosakti/kitchen-grpc/services/orders/service"
)

type httpServer struct {
	addr string
}

func newHTTPOrderServer(addr string) *httpServer {
	return &httpServer{
		addr: addr,
	}
}

func (s *httpServer) Run() error {
	router := http.NewServeMux()

	orderService := service.NewOrderService()
	orderHandler := handler.NewHTTPOrderHandler(orderService)
	orderHandler.RegisterRoute(router)

	log.Println("Starting http server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
