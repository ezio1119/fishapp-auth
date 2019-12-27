package interactor

import (
	"time"

	"github.com/ezio1119/fishapp-auth/usecase/repository"
)

type AuthInteractor struct {
	AuthRepository  repository.AuthRepository
	TokenInteractor UTokenInteractor
	ContextTimeout  time.Duration
}

type UAuthInteractor interface {
	AddBlackList(t string) error
	CheckBlackList(t string) (bool, error)
}

func (i *AuthInteractor) AddBlackList(t string) error {
	return i.AuthRepository.SAdd(t)
}

func (i *AuthInteractor) CheckBlackList(t string) (bool, error) {
	return i.AuthRepository.SIsMember(t)
}
