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
}
