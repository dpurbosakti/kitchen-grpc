package handler

import (
	"context"

	"github.com/dpurbosakti/kitchen-grpc/services/common/genproto/orders"
	"github.com/dpurbosakti/kitchen-grpc/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGRPCHandler struct {
	orderService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGRPCOrdersService(gRPC *grpc.Server, ordersService types.OrderService) {
	gRPCHandler := &OrdersGRPCHandler{
		orderService: ordersService,
	}

	orders.RegisterOrderServiceServer(gRPC, gRPCHandler)
}

func (h *OrdersGRPCHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    42,
		CustomerID: 2,
		ProductID:  1,
		Quantity:   10,
	}

	err := h.orderService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	return res, nil
}

func (h *OrdersGRPCHandler) GetOrders(ctx context.Context, req *orders.GetOrdersRequest) (*orders.GetOrdersResponse, error) {
	o, err := h.orderService.GetOrders(ctx)
	if err != nil {
		return nil, err
	}

	res := &orders.GetOrdersResponse{
		Orders: o,
	}

	return res, nil
}
