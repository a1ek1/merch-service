package usecase

import (
	"merch-service/internal/domain/service"
)

type CoinUsecase interface {
	SendCoins(fromUsername, toUsername string, amount int) error
}

type coinUsecase struct {
	coinService service.CoinService
}

func NewCoinUsecase(coinService service.CoinService) CoinUsecase {
	return &coinUsecase{
		coinService: coinService,
	}
}

func (c *coinUsecase) SendCoins(fromUsername, toUsername string, amount int) error {
	return c.coinService.SendCoin(fromUsername, toUsername, amount)
}
