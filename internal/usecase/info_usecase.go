package usecase

import (
	"github.com/google/uuid"
	"merch-service/internal/domain/service"
)

type InfoUsecase interface {
	GetUserInfo(userID uuid.UUID) (*service.InfoResponse, error)
}

type infoUsecase struct {
	infoService service.InfoService
}

func NewInfoUsecase(infoService service.InfoService) InfoUsecase {
	return &infoUsecase{
		infoService: infoService,
	}
}

func (i *infoUsecase) GetUserInfo(userID uuid.UUID) (*service.InfoResponse, error) {
	return i.infoService.GetUserInfo(userID)
}
