package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"log"
	"merch-service/internal/interactor"
	"merch-service/internal/presenter/http/middleware"
	"merch-service/internal/presenter/http/router"
	"merch-service/pkg/config"
	"time"
)

func main() {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Get().DBHost, config.Get().DBPort, config.Get().DBUser,
		config.Get().DBPassword, config.Get().DBName, config.Get().SSLMode,
	)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

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

	log.Fatal(e.Start(":8080"))
}
