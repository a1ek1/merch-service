package usecase

import (
	"github.com/google/uuid"
	"merch-service/internal/domain/service"
)

type PurchaseUsecase interface {
	BuyItem(userID uuid.UUID, itemName string) error
}

type purchaseUsecase struct {
	purchaseService service.PurchaseService
}

func NewPurchaseUsecase(purchaseService service.PurchaseService) PurchaseUsecase {
	return &purchaseUsecase{
		purchaseService: purchaseService,
	}
}

func (p *purchaseUsecase) BuyItem(userID uuid.UUID, itemName string) error {
	return p.purchaseService.BuyItem(userID, itemName)
}
