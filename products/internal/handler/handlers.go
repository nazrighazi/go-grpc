package handler

import (
	service "go_grpc/products/internal/service"
	pb "go_grpc/proto/products"
)
type ProductHandlerDto struct {
	Service service.IService
}
type ProductHandler struct {
	pb.UnimplementedProductsServer
	service service.IService
}

func NewProductHandler(req ProductHandlerDto) *ProductHandler {
	return &ProductHandler{service: req.Service}
}