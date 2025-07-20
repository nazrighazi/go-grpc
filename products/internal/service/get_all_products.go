package service

import (
	"context"
	"database/sql"
	pb "go_grpc/proto/products"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetAllProducts(ctx context.Context) (*pb.GetAllProductsResponse, error) {

	products, err := s.dataAccess.AllProducts(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.GetAllProductsResponse{Data: nil, Status: "success", Message: "no products found"}, nil
		}
		s.log.Error("failed to get all products from database", "error", err)
		return nil, status.Errorf(codes.Internal, "failed to get products: %v", err)
	}

	if len(products) == 0 {
		return &pb.GetAllProductsResponse{Data: nil, Status: "success", Message: "no products found"}, nil
	}

	var response []*pb.Product

	for _, product := range products {

		price, err := product.Price.Float64Value()
		if err != nil {
			s.log.Error("failed to convert product price", "error", err)
			return nil, status.Errorf(codes.Internal, "failed to process product data: %v", err)
		}
		response = append(response, &pb.Product{
			Name:        product.Name,
			Price:       price.Float64,
			Description: product.Description.String,
		})
	}
	return &pb.GetAllProductsResponse{Data: response, Status: "success", Message: "products found"}, nil
}

