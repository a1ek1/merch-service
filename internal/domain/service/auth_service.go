package service

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"merch-service/internal/domain/model"
	"merch-service/internal/domain/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Authenticate(username, password string) (string, error)
	ValidateToken(tokenStr string) (*jwt.Token, error)
}

type authService struct {
	userRepo  repository.UserRepository
	tokenRepo repository.AuthTokenRepository
	jwtSecret string
}

func (a authService) Authenticate(username, password string) (string, error) {
	user, err := a.userRepo.GetUserByUsername(username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			passwordHash, err := hashPassword(password)
			if err != nil {
				return "", err
			}
			user = &model.User{
				ID:       uuid.New(),
				Username: username,
				Password: passwordHash,
			}
			if err := a.userRepo.Create(user); err != nil {
				return "", errors.New("error creating user")
			}
		} else {
			return "", err
		}
	} else {
		if err := comparePasswords(user.Password, password); err != nil {
			return "", errors.New("incorrect password")
		}
	}

	existingToken, err := a.tokenRepo.GetActiveTokenByUserID(user.ID)
	if err == nil && existingToken != nil && existingToken.ExpiresAt.After(time.Now()) {
		return existingToken.Token, nil
	}

	tokenStr, err := a.generateToken(user)
	if err != nil {
		return "", err
	}

	authToken := &model.AuthToken{
		UserID:    user.ID,
		Token:     tokenStr,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	if err := a.tokenRepo.Create(authToken); err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (a authService) generateToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(a.jwtSecret))
}

func (a authService) ValidateToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func NewAuthService(userRepo repository.UserRepository, tokenRepo repository.AuthTokenRepository, jwtSecret string) AuthService {
	return &authService{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
		jwtSecret: jwtSecret,
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func comparePasswords(hashedPwd string, plainPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
}
