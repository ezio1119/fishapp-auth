package user

import (
	"context"

	"github.com/ezio1119/fishapp-user/models"
)

// Usecase represent the user's usecases
type Usecase interface {
	GetByID(ctx context.Context, id int64) (*models.User, error)
	Update(ctx context.Context, u *models.User) error
	Create(ctx context.Context, u *models.User) (*models.TokenPair, error)
	Delete(ctx context.Context, id int64) error
	RefreshIDToken(ctx context.Context, id int64) (*models.TokenPair, error)
	Login(ctx context.Context, email string, pass string) (*models.User, *models.TokenPair, error)
}
