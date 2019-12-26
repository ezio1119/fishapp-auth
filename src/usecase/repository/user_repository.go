package repository

import (
	"context"

	"github.com/ezio1119/fishapp-auth/domain"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	Update(ctx context.Context, u *domain.User) error
	Create(ctx context.Context, u *domain.User) error
	Delete(ctx context.Context, id int64) error
}
