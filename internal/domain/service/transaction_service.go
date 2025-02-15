package service

import (
	"errors"
	"fmt"
	"time"

	"merch-service/internal/domain/model"
	"merch-service/internal/domain/repository"

	"github.com/google/uuid"
)

type CoinService interface {
	SendCoin(fromUsername, toUsername string, amount int) error
}

type coinService struct {
	userRepo        repository.UserRepository
	transactionRepo repository.TransactionRepository
}

func NewCoinService(userRepo repository.UserRepository, transactionRepo repository.TransactionRepository) CoinService {
	return &coinService{
		userRepo:        userRepo,
		transactionRepo: transactionRepo,
	}
}

func (s *coinService) SendCoin(fromUsername, toUsername string, amount int) error {
	if amount <= 0 {
		return errors.New("сумма перевода должна быть положительной")
	}

	fromUser, err := s.userRepo.GetUserByUsername(fromUsername)
	if err != nil {
		return fmt.Errorf("ошибка при получении отправителя: %w", err)
	}
	if fromUser == nil {
		return errors.New("отправитель не найден")
	}

	toUser, err := s.userRepo.GetUserByUsername(toUsername)
	if err != nil {
		return fmt.Errorf("ошибка при получении получателя: %w", err)
	}
	if toUser == nil {
		return errors.New("получатель не найден")
	}

	if fromUser.Coins < amount {
		return errors.New("недостаточно монет для перевода")
	}

	newSenderBalance := fromUser.Coins - amount
	newReceiverBalance := toUser.Coins + amount

	if err := s.userRepo.UpdateBalance(fromUser.ID, newSenderBalance); err != nil {
		return fmt.Errorf("ошибка при обновлении баланса отправителя: %w", err)
	}

	if err := s.userRepo.UpdateBalance(toUser.ID, newReceiverBalance); err != nil {
		revertErr := s.userRepo.UpdateBalance(fromUser.ID, fromUser.Coins)
		if revertErr != nil {
			return fmt.Errorf("ошибка при откате обновления баланса отправителя после сбоя: %w", revertErr)
		}
		return fmt.Errorf("ошибка при обновлении баланса получателя: %w", err)
	}

	transaction := &model.Transaction{
		ID:        uuid.New(),
		UserID:    fromUser.ID,
		ToUserID:  toUser.ID,
		Amount:    amount,
		CreatedAt: time.Now(),
	}

	if err := s.transactionRepo.Create(transaction); err != nil {
		revertErr := s.userRepo.UpdateBalance(fromUser.ID, fromUser.Coins)
		if revertErr != nil {
			return fmt.Errorf("ошибка при откате обновления баланса отправителя после сбоя транзакции: %w", revertErr)
		}
		revertErr = s.userRepo.UpdateBalance(toUser.ID, toUser.Coins)
		if revertErr != nil {
			return fmt.Errorf("ошибка при откате обновления баланса получателя после сбоя транзакции: %w", revertErr)
		}
		return fmt.Errorf("ошибка при создании транзакции: %w", err)
	}

	return nil
}
