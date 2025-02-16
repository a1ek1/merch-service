package handler

import (
	"github.com/labstack/echo"
	"merch-service/internal/domain/model"
	"merch-service/internal/usecase"
	"net/http"
)

// AuthHandler интерфейс для обработки аутентификации
type AuthHandler interface {
	Authenticate(c echo.Context) error
}

// authHandler - структура для обработчика аутентификации
type authHandler struct {
	authUsecase usecase.AuthUsecase
}

// NewAuthHandler - создаёт новый обработчик для аутентификации
func NewAuthHandler(authUsecase usecase.AuthUsecase) AuthHandler {
	return &authHandler{
		authUsecase: authUsecase,
	}
}

// Authenticate - обработчик аутентификации
func (h *authHandler) Authenticate(c echo.Context) error {
	var request model.AuthRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	token, err := h.authUsecase.AuthenticateUser(request.Username, request.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authentication failed")
	}

	return c.JSON(http.StatusOK, model.AuthResponse{
		Token: token,
	})
}
