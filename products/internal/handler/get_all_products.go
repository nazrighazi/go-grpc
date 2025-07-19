package handler

import (
	"context"
	"errors"
	pb "go_grpc/proto/products"
)

func (handler *ProductHandler) GetAllProducts(ctx context.Context, req *pb.GetAllProductsRequest) (*pb.GetAllProductsResponse, error) {

	products, err := handler.service.GetAllProducts(ctx)
	if err != nil {
		return nil, errors.New("failed to get all products")
	}

	var response []*pb.Product

	if len(products) == 0 {
		return &pb.GetAllProductsResponse{Data: nil, Status: "success", Message: "no products found"}, nil
	}

	for _, product := range products {

		price, err := product.Price.Float64Value()
		if err != nil {
			// Handle the error, for example:
			return nil, err
		}
		response = append(response, &pb.Product{
			Name:        product.Name,
			Price:       price.Float64,
			Description: product.Description.String,
		})
	}
	return &pb.GetAllProductsResponse{Data: response, Status: "success", Message: "products found"}, nil
}