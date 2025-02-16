package model

import "github.com/dgrijalva/jwt-go"

// AuthRequest - модель для запроса аутентификации
type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthResponse - модель для ответа на запрос аутентификации (токен)
type AuthResponse struct {
	Token string `json:"token"`
}

// SendCoinRequest - модель для запроса на отправку монет
type SendCoinRequest struct {
	FromUsername string `json:"fromUsername"`
	ToUsername   string `json:"toUsername"`
	Amount       int    `json:"amount"`
}

// InfoResponse - модель для ответа с информацией о пользователе
type InfoResponse struct {
	Coins       int             `json:"coins"`
	Inventory   []InventoryItem `json:"inventory"`
	CoinHistory CoinHistory     `json:"coinHistory"`
}

// InventoryItem - модель для элемента инвентаря
type InventoryItem struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

// CoinHistory - история монет (полученные и отправленные)
type CoinHistory struct {
	Received []CoinTransaction `json:"received"`
	Sent     []CoinTransaction `json:"sent"`
}

// CoinTransaction - модель для транзакции монет
type CoinTransaction struct {
	FromUser string `json:"fromUser,omitempty"`
	ToUser   string `json:"toUser,omitempty"`
	Amount   int    `json:"amount"`
}

// ErrorResponse - модель для ответа с ошибкой
type ErrorResponse struct {
	Errors string `json:"errors"`
}

// JWTClaims - модель для парсинга и формирования jwt токена
type JWTClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
