package usecase

import (
	"merch-service/internal/domain/service"
)

type AuthUsecase interface {
	AuthenticateUser(username, password string) (string, error)
}

type authUsecase struct {
	authService service.AuthService
}

func NewAuthUsecase(authService service.AuthService) AuthUsecase {
	return &authUsecase{
		authService: authService,
	}
}

func (a *authUsecase) AuthenticateUser(username, password string) (string, error) {
	return a.authService.Authenticate(username, password)
}
