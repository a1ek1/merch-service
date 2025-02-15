package service

import (
	"errors"
	"time"

	"merch-service/internal/domain/model"
	"merch-service/internal/domain/repository"

	"github.com/google/uuid"
)

type PurchaseService interface {
	BuyItem(userID uuid.UUID, itemName string) error
}

type purchaseService struct {
	userRepo      repository.UserRepository
	itemRepo      repository.ItemRepository
	purchaseRepo  repository.PurchaseRepository
	inventoryRepo repository.InventoryRepository
}

func NewPurchaseService(
	userRepo repository.UserRepository,
	itemRepo repository.ItemRepository,
	purchaseRepo repository.PurchaseRepository,
	inventoryRepo repository.InventoryRepository,
) PurchaseService {
	return &purchaseService{
		userRepo:      userRepo,
		itemRepo:      itemRepo,
		purchaseRepo:  purchaseRepo,
		inventoryRepo: inventoryRepo,
	}
}

func (s *purchaseService) BuyItem(userID uuid.UUID, itemName string) error {
	item, err := s.itemRepo.GetItemByTitle(itemName)
	if err != nil {
		return err
	}

	user, err := s.userRepo.GetUserById(userID)
	if err != nil {
		return err
	}

	if user.Coins < item.Price {
		return errors.New("недостаточно монет для покупки")
	}

	newBalance := user.Coins - item.Price
	if err := s.userRepo.UpdateBalance(user.ID, newBalance); err != nil {
		return err
	}

	purchase := &model.Purchase{
		ID:          uuid.New(),
		UserID:      user.ID,
		ItemID:      item.ID,
		Quantity:    1,
		TotalPrice:  item.Price,
		PurchasedAt: time.Now(),
	}
	if err := s.purchaseRepo.Create(purchase); err != nil {
		return err
	}

	if err := s.inventoryRepo.AddItem(user.ID, item.ID, 1); err != nil {
		return err
	}

	return nil
}
