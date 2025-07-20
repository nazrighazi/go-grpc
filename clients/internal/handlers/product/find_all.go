package product_handler

import (
	pb "go_grpc/proto/products"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (handler *ProductHandler) FindAll(c echo.Context) error {

	products, err := handler.grpcServer.Product.GetAllProducts(c.Request().Context(), &pb.GetAllProductsRequest{})
	if err != nil {
		handler.log.Error("Failed to get products from gRPC server", "error", err)
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "error",
			"message": "failed to get products",
			"data":    nil,
			"error": map[string]any{
				"message": err.Error(),
			},
		})
	}

	handler.log.Debug("Received products from gRPC server", "products", products)

	responseData := map[string]any{
		"status":  products.Status,
		"message": products.Message,
		"data":    products.Data,
	}

	if len(products.Data) == 0 {
		responseData["data"] = nil
	}

	return c.JSON(http.StatusOK, responseData)
}