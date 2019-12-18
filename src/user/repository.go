package user

import (
	"context"

	"github.com/ezio1119/fishapp-user/models"
)

// Repository represent the user's repository contract
type Repository interface {
	GetByID(ctx context.Context, id int64) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Update(ctx context.Context, u *models.User) error
	Create(ctx context.Context, u *models.User) error
	Delete(ctx context.Context, id int64) error
}
