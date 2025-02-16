//package main
//
//import (
//	"fmt"
//	"log"
//
//	"github.com/jmoiron/sqlx"
//	_ "github.com/lib/pq"
//	"merch-service/internal/domain/service"
//	"merch-service/internal/infrastructure/datastore"
//)
//
//func main() {
//	// Подключение к базе данных
//	connStr := "host=localhost port=5435 user=postgres password=postgres dbname=merch_service sslmode=disable"
//	db, err := sqlx.Connect("postgres", connStr)
//	if err != nil {
//		log.Fatal("Ошибка подключения к базе данных: ", err)
//	}
//
//	// Создание репозиториев
//	userRepo := datastore.NewUserRepository(db)
//	itemRepo := datastore.NewItemRepository(db)
//	purchaseRepo := datastore.NewPurchaseRepository(db)
//	inventoryRepo := datastore.NewInventoryRepository(db)
//	authTokenRepo := datastore.NewAuthTokenRepository(db)
//	transactionRepo := datastore.NewTransactionRepository(db)
//
//	// Создаем сервисы
//	jwtSecret := "your-jwt-secret" // Укажите свой секретный ключ для JWT
//	authService := service.NewAuthService(userRepo, authTokenRepo, jwtSecret)
//	purchaseService := service.NewPurchaseService(userRepo, itemRepo, purchaseRepo, inventoryRepo)
//
//	// Создаем сервис для работы с монетами
//	coinService := service.NewCoinService(userRepo, transactionRepo)
//
//	// Тестовый пользователь и товар
//	username := "testuser10"   // Имя существующего пользователя
//	password := "testpassword" // Пароль пользователя
//
//	// Аутентификация
//	token, err := authService.Authenticate(username, password)
//	if err != nil {
//		log.Fatal("Ошибка аутентификации: ", err)
//	}
//
//	fmt.Printf("Полученный токен: %s\n", token)
//
//	// Валидация токена
//	validatedToken, err := authService.ValidateToken(token)
//	if err != nil {
//		log.Fatal("Ошибка валидации токена: ", err)
//	}
//
//	fmt.Printf("Валидированный токен: %+v\n", validatedToken)
//
//	// Получаем ID пользователя по токену
//	authToken, err := authTokenRepo.GetByToken(token)
//	if err != nil {
//		log.Fatal("Ошибка получения токена из репозитория: ", err)
//	}
//
//	// Получаем ID пользователя из токена
//	userID := authToken.UserID
//	fmt.Printf("ID пользователя: %s\n", userID)
//
//	// Попытка покупки товара
//	itemName := "t-shirt" // Название товара для покупки
//
//	err = purchaseService.BuyItem(userID, itemName)
//	if err != nil {
//		log.Fatalf("Ошибка при покупке товара: %v", err)
//	}
//
//	// Получаем обновленный баланс пользователя
//	user, err := userRepo.GetUserById(userID)
//	if err != nil {
//		log.Fatalf("Ошибка при получении данных пользователя: %v", err)
//	}
//
//	// Выводим новый баланс
//	fmt.Printf("Новый баланс пользователя %s: %d монет\n", user.Username, user.Coins)
//
//	// Выводим информацию о инвентаре пользователя
//	inventory, err := inventoryRepo.GetByUserID(userID)
//	if err != nil {
//		log.Fatalf("Ошибка при получении инвентаря пользователя: %v", err)
//	}
//
//	// Выводим инвентарь
//	fmt.Println("Инвентарь пользователя:")
//	for _, inv := range inventory {
//		fmt.Printf("- Товар: %s, Количество: %d\n", inv.ItemID.String(), inv.Quantity)
//	}
//
//	// Пример перевода монет между пользователями
//	fromUsername := "testuser10" // Отправитель
//	toUsername := "testuser2"    // Получатель
//	amount := 30                 // Сумма перевода
//
//	// Перевод монет
//	err = coinService.SendCoin(fromUsername, toUsername, amount)
//	if err != nil {
//		log.Printf("Ошибка при переводе монет: %v", err)
//	} else {
//		fmt.Printf("Перевод %d монет от %s к %s выполнен успешно\n", amount, fromUsername, toUsername)
//	}
//}

package main

import (
	"fmt"
	"log"
	"merch-service/config"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"

	"merch-service/internal/interactor"
	"merch-service/internal/presenter/http/middleware"
	"merch-service/internal/presenter/http/router"
)

func main() {
	dbHost := config.Get().DBHost
	dbPort := config.Get().DBPort
	dbUser := config.Get().DBUser
	dbPassword := config.Get().DBPassword
	dbName := config.Get().DBName
	sslMode := config.Get().SSLMode

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Failed to close database connection: %v", err)
		}
	}()

	i := interactor.NewInteractor(db)

	h := i.NewAppHandler()

	e := echo.New()

	router.NewRouter(e, h)
	middleware.NewMiddleware(e)

	port := config.Get().APPPort

	log.Printf("Starting server on port %s", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
