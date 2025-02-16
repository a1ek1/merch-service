package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"merch-service/internal/usecase"
	"net/http"
)

// PurchaseHandler интерфейс для обработки покупок предметов
type PurchaseHandler interface {
	BuyItem(c echo.Context) error
}

// purchaseHandler - структура для обработчика покупок
type purchaseHandler struct {
	purchaseUsecase usecase.PurchaseUsecase
}

// NewPurchaseHandler - создаёт новый обработчик для покупок
func NewPurchaseHandler(purchaseUsecase usecase.PurchaseUsecase) PurchaseHandler {
	return &purchaseHandler{
		purchaseUsecase: purchaseUsecase,
	}
}

// BuyItem - обработчик покупки предмета
func (h *purchaseHandler) BuyItem(c echo.Context) error {
	itemName := c.Param("item")

	userID, ok := c.Get("userID").(string)
	if !ok || userID == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid or missing user ID"})
	}

	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid user ID format"})
	}

	err = h.purchaseUsecase.BuyItem(parsedUserID, itemName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to buy item")
	}

	return c.NoContent(http.StatusOK)
}
