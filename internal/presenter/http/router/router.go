package router

import (
	"github.com/labstack/echo"
	"merch-service/internal/presenter/http/handler"
)

// NewRouter sets up routes for the transaction service.
func NewRouter(e *echo.Echo, h handler.AppHandler) {
	api := e.Group("/api")
	{
		api.POST("/sendCoin", h.SendCoins)
		api.GET("/info", h.GetUserInfo)
		api.POST("/auth", h.Authenticate)
		api.GET("/buy/:item", h.BuyItem)
	}
}
