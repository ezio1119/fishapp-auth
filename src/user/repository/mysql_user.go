package repository

import (
	"context"

	"github.com/ezio1119/fishapp-user/models"
	"github.com/ezio1119/fishapp-user/user"
	"github.com/jinzhu/gorm"
)

type mysqlUserRepository struct {
	Conn *gorm.DB
}

// NewMysqluserRepository will create an object that represent the user.Repository interface
func NewMysqlUserRepository(Conn *gorm.DB) user.Repository {
	return &mysqlUserRepository{Conn}
}

func (m *mysqlUserRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	var user models.User
	if err := m.Conn.Where("id = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *mysqlUserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := m.Conn.Where("email = ?", email).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *mysqlUserRepository) Create(ctx context.Context, u *models.User) error {
	if err := m.Conn.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) Update(ctx context.Context, u *models.User) error {
	var user models.User
	if err := m.Conn.Where("id = ?", u.ID).Take(&user).Updates(u).Error; err != nil {
		return err
	}
	u.UpdatedAt = user.UpdatedAt
	u.CreatedAt = user.CreatedAt
	return nil
}

func (m *mysqlUserRepository) Delete(ctx context.Context, id int64) error {
	if err := m.Conn.Where("id = ?", id).Take(&models.User{}).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}
