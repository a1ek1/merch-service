package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"merch-service/internal/usecase"
	"net/http"
)

// InfoHandler интерфейс для обработки запросов, связанных с информацией о пользователе
type InfoHandler interface {
	GetUserInfo(c echo.Context) error
}

// infoHandler - структура для обработчика информации о пользователе
type infoHandler struct {
	infoUsecase usecase.InfoUsecase
}

// NewInfoHandler - создаёт новый обработчик для получения информации о пользователе
func NewInfoHandler(infoUsecase usecase.InfoUsecase) InfoHandler {
	return &infoHandler{
		infoUsecase: infoUsecase,
	}
}

// GetUserInfo - обработчик получения информации о пользователе
func (h *infoHandler) GetUserInfo(c echo.Context) error {
	// Теперь userID приходит уже в формате UUID из middleware.go
	userID, ok := c.Get("userID").(string)
	if !ok || userID == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid or missing user ID"})
	}

	// Преобразуем строку в UUID
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid user ID format"})
	}

	info, err := h.infoUsecase.GetUserInfo(parsedUserID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user info")
	}

	return c.JSON(http.StatusOK, info)
}
