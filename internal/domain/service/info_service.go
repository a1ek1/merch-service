package service

import (
	"github.com/google/uuid"
	"merch-service/internal/domain/repository"
)

type InfoResponse struct {
	Coins       int                     `json:"coins"`
	Inventory   []InventoryItemResponse `json:"inventory"`
	CoinHistory CoinHistoryResponse     `json:"coinHistory"`
}

type InventoryItemResponse struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

type TransactionHistory struct {
	FromUser string `json:"fromUser,omitempty"`
	ToUser   string `json:"toUser,omitempty"`
	Amount   int    `json:"amount"`
}

type CoinHistoryResponse struct {
	Received []TransactionHistory `json:"received"`
	Sent     []TransactionHistory `json:"sent"`
}

type InfoService interface {
	GetUserInfo(userID uuid.UUID) (*InfoResponse, error)
}

type infoService struct {
	userRepo        repository.UserRepository
	inventoryRepo   repository.InventoryRepository
	transactionRepo repository.TransactionRepository
}

func NewInfoService(
	userRepo repository.UserRepository,
	inventoryRepo repository.InventoryRepository,
	transactionRepo repository.TransactionRepository,
) InfoService {
	return &infoService{
		userRepo:        userRepo,
		inventoryRepo:   inventoryRepo,
		transactionRepo: transactionRepo,
	}
}

func (s *infoService) GetUserInfo(userID uuid.UUID) (*InfoResponse, error) {
	user, err := s.userRepo.GetUserById(userID)
	if err != nil {
		return nil, err
	}

	inventory, err := s.inventoryRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	var inventoryResp []InventoryItemResponse
	for _, inv := range inventory {
		inventoryResp = append(inventoryResp, InventoryItemResponse{
			Type:     inv.ItemID.String(),
			Quantity: inv.Quantity,
		})
	}

	transactions, err := s.transactionRepo.GetByUserId(userID)
	if err != nil {
		return nil, err
	}
	var sent, received []TransactionHistory
	for _, tx := range transactions {
		if tx.UserID == userID {
			sent = append(sent, TransactionHistory{
				ToUser: tx.ToUserID.String(),
				Amount: tx.Amount,
			})
		}
		if tx.ToUserID == userID {
			received = append(received, TransactionHistory{
				FromUser: tx.UserID.String(),
				Amount:   tx.Amount,
			})
		}
	}

	coinHistory := CoinHistoryResponse{
		Received: received,
		Sent:     sent,
	}

	return &InfoResponse{
		Coins:       user.Coins,
		Inventory:   inventoryResp,
		CoinHistory: coinHistory,
	}, nil
}
