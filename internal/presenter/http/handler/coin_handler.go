package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"merch-service/internal/domain/model"
	"merch-service/internal/usecase"
	"net/http"
)

// CoinHandler интерфейс для обработки транзакций с монетами
type CoinHandler interface {
	SendCoins(c echo.Context) error
}

// coinHandler - структура для обработчика транзакций с монетами
type coinHandler struct {
	coinUsecase usecase.CoinUsecase
}

// SendCoins - обработчик отправки монет
func (h *coinHandler) SendCoins(c echo.Context) error {
	username, ok := c.Get("username").(string)
	if !ok || username == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}

	var request model.SendCoinRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	fmt.Println(username, request.ToUsername, request.Amount)
	err := h.coinUsecase.SendCoins(username, request.ToUsername, request.Amount)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to send coins")
	}

	return c.NoContent(http.StatusOK)
}

// NewCoinHandler - создаёт новый обработчик для транзакций с монетами
func NewCoinHandler(coinUsecase usecase.CoinUsecase) CoinHandler {
	return &coinHandler{
		coinUsecase: coinUsecase,
	}
}
