package handler

type AppHandler interface {
	AuthHandler
	CoinHandler
	InfoHandler
	PurchaseHandler
}
