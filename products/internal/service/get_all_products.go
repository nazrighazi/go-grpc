package service

import (
	"context"
	"errors"
	dataaccess "go_grpc/products/internal/data_access"
)

func (s *Service) GetAllProducts(ctx context.Context) ([]dataaccess.Product, error) {

	products, err := s.dataAccess.AllProducts(ctx)
	if err != nil {
		return nil, errors.New("failed to get all products")
	}
	return products, nil
}