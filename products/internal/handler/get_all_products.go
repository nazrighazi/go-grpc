package handler

import (
	"context"
	"errors"
	pb "go_grpc/proto/products"
)

func (handler *ProductHandler) GetAllProducts(ctx context.Context, req *pb.GetAllProductsRequest) (*pb.GetAllProductsResponse, error) {

	response, err := handler.service.GetAllProducts(ctx)
	if err != nil {
		return nil, errors.New("failed to get all products")
	}

	return response, nil
}