package products_handler

import (
	pb "go_grpc/proto/products"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (handler *ProductsHandler) FindAll(c echo.Context) error {

	products, err := handler.grpcServer.GetAllProducts(c.Request().Context(), &pb.GetAllProductsRequest{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, products)
}