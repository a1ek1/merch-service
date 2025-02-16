package interactor

import (
	"github.com/jmoiron/sqlx"
	"merch-service/internal/domain/repository"
	"merch-service/internal/domain/service"
	"merch-service/internal/infrastructure/datastore"
	"merch-service/internal/presenter/http/handler"
	"merch-service/internal/usecase"
)

type Interactor interface {
	NewAuthTokenRepository() repository.AuthTokenRepository
	NewInventoryRepository() repository.InventoryRepository
	NewItemRepository() repository.ItemRepository
	NewPurchaseRepository() repository.PurchaseRepository
	NewTransactionRepository() repository.TransactionRepository
	NewUserRepository() repository.UserRepository
	NewAuthService() service.AuthService
	NewInfoService() service.InfoService
	NewPurchaseService() service.PurchaseService
	NewCoinService() service.CoinService
	NewAuthUsecase() usecase.AuthUsecase
	NewCoinUsecase() usecase.CoinUsecase
	NewInfoUsecase() usecase.InfoUsecase
	NewPurchaseUsecase() usecase.PurchaseUsecase
	NewAuthHandler() handler.AuthHandler
	NewCoinHandler() handler.CoinHandler
	NewInfoHandler() handler.InfoHandler
	NewPurchaseHandler() handler.PurchaseHandler
	NewAppHandler() handler.AppHandler
}
type interactor struct {
	DB *sqlx.DB
}

func (i *interactor) NewAuthTokenRepository() repository.AuthTokenRepository {
	return datastore.NewAuthTokenRepository(i.DB)
}

func (i *interactor) NewInventoryRepository() repository.InventoryRepository {
	return datastore.NewInventoryRepository(i.DB)
}

func (i *interactor) NewItemRepository() repository.ItemRepository {
	return datastore.NewItemRepository(i.DB)
}

func (i *interactor) NewPurchaseRepository() repository.PurchaseRepository {
	return datastore.NewPurchaseRepository(i.DB)
}

func (i *interactor) NewTransactionRepository() repository.TransactionRepository {
	return datastore.NewTransactionRepository(i.DB)
}

func (i *interactor) NewUserRepository() repository.UserRepository {
	return datastore.NewUserRepository(i.DB)
}

func (i *interactor) NewAuthService() service.AuthService {
	return service.NewAuthService(i.NewUserRepository(), i.NewAuthTokenRepository(), "my_key")
}

func (i *interactor) NewInfoService() service.InfoService {
	return service.NewInfoService(i.NewUserRepository(), i.NewInventoryRepository(), i.NewTransactionRepository())
}

func (i *interactor) NewPurchaseService() service.PurchaseService {
	return service.NewPurchaseService(i.NewUserRepository(), i.NewItemRepository(), i.NewPurchaseRepository(), i.NewInventoryRepository())
}

func (i *interactor) NewCoinService() service.CoinService {
	return service.NewCoinService(i.NewUserRepository(), i.NewTransactionRepository())
}

func (i *interactor) NewCoinUsecase() usecase.CoinUsecase {
	return usecase.NewCoinUsecase(i.NewCoinService())
}

func (i *interactor) NewInfoUsecase() usecase.InfoUsecase {
	return usecase.NewInfoUsecase(i.NewInfoService())
}

func (i *interactor) NewPurchaseUsecase() usecase.PurchaseUsecase {
	return usecase.NewPurchaseUsecase(i.NewPurchaseService())
}

func (i *interactor) NewAuthUsecase() usecase.AuthUsecase {
	return usecase.NewAuthUsecase(i.NewAuthService())
}

func (i *interactor) NewAuthHandler() handler.AuthHandler {
	return handler.NewAuthHandler(i.NewAuthUsecase())
}

func (i *interactor) NewCoinHandler() handler.CoinHandler {
	return handler.NewCoinHandler(i.NewCoinUsecase())
}

func (i *interactor) NewInfoHandler() handler.InfoHandler {
	return handler.NewInfoHandler(i.NewInfoUsecase())
}

func (i *interactor) NewPurchaseHandler() handler.PurchaseHandler {
	return handler.NewPurchaseHandler(i.NewPurchaseUsecase())
}

func NewInteractor(db *sqlx.DB) Interactor {
	return &interactor{DB: db}
}

type appHandler struct {
	handler.CoinHandler
	handler.AuthHandler
	handler.InfoHandler
	handler.PurchaseHandler
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	return &appHandler{
		i.NewCoinHandler(),
		i.NewAuthHandler(),
		i.NewInfoHandler(),
		i.NewPurchaseHandler(),
	}
}
