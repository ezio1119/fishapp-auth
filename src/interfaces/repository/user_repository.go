package repository

import (
	"context"

	"github.com/ezio1119/fishapp-auth/domain"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	var user domain.User
	if err := r.Conn.Where("id = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := r.Conn.Where("email = ?", email).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, u *domain.User) error {
	if err := r.Conn.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Update(ctx context.Context, u *domain.User) error {
	var user domain.User
	if err := r.Conn.Where("id = ?", u.ID).Take(&user).Updates(u).Error; err != nil {
		return err
	}
	u.UpdatedAt = user.UpdatedAt
	u.CreatedAt = user.CreatedAt
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	if err := r.Conn.Where("id = ?", id).Take(&domain.User{}).Delete(&domain.User{}).Error; err != nil {
		return err
	}
	return nil
}
