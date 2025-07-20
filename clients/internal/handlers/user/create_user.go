package user_handler

import (
	"fmt"
	users_proto "go_grpc/proto/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	CreateUserRequestDto struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	CreateUserResponseDto struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Data    any    `json:"data"`
	}
)
func (handler *UserHandler) CreateUser(c echo.Context) (error) {

	requestBody := new(CreateUserRequestDto)
	if err := c.Bind(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "invalid request body",
			"data":    nil,
		})
	}

	response, err := handler.grpcServer.User.RegisterUser(c.Request().Context(), &users_proto.RegisterUserRequestDto{Email:requestBody.Email,Password: requestBody.Password}); 
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"status":  "error",
				"message": "failed to create user",
				"data":    nil,
			})
		}

	return c.JSON(http.StatusCreated, CreateUserResponseDto{
		Status:  response.Status,
		Message: response.Message,
		Data:    nil,
	})
}