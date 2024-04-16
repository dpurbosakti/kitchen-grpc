package handler

import (
	"net/http"

	"github.com/dpurbosakti/kitchen-grpc/services/common/genproto/orders"
	"github.com/dpurbosakti/kitchen-grpc/services/common/util"
	"github.com/dpurbosakti/kitchen-grpc/services/orders/types"
)

type OrderHTTPHandler struct {
	ordersService types.OrderService
}

func NewHTTPOrderHandler(orderService types.OrderService) *OrderHTTPHandler {
	return &OrderHTTPHandler{
		ordersService: orderService,
	}
}

func (h *OrderHTTPHandler) RegisterRoute(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrderHTTPHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &orders.Order{
		OrderID:    42,
		CustomerID: req.GetCustomerID(),
		ProductID:  req.GetProductID(),
		Quantity:   req.GetQuantity(),
	}

	err = h.ordersService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}
	util.WriteJSON(w, http.StatusOK, res)

}
