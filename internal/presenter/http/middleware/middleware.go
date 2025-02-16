package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Secret key for JWT signing and validation
var jwtSecret = []byte("my_key")

// Claims struct for JWT
type Claims struct {
	UserID   string `json:"user_id"` // Добавлен UserID
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken creates a JWT token for the user
func GenerateToken(userID, username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token valid for 24 hours
	claims := &Claims{
		UserID:   userID, // Теперь включаем UserID
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// JWTMiddleware verifies the token and extracts user info
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Пропускаем авторизацию для /api/auth
		if c.Path() == "/api/auth" {
			return next(c)
		}

		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Authorization header is required"})
		}

		// Token should be in "Bearer {token}" format
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token format"})
		}

		tokenString := tokenParts[1]
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid or expired token"})
		}

		// Логируем, что у нас есть в токене
		fmt.Println("Extracted from token - UserID:", claims.UserID, "Username:", claims.Username)

		c.Set("userID", claims.UserID) // Сохраняем userID
		c.Set("username", claims.Username)

		return next(c)
	}
}

// NewMiddleware applies global middlewares to the Echo instance
func NewMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.Use(JWTMiddleware) // Применяем исправленный JWT миддлвар
}
