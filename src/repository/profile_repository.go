package repository

import (
	"context"

	"github.com/ezio1119/fishapp-profile/domain"
	"github.com/jinzhu/gorm"
)

// Usecase
type ProfileRepository interface {
	GetByID(ctx context.Context, userID int64) (*domain.Profile, error)
	Update(ctx context.Context, p *domain.Profile) error
	Create(ctx context.Context, p *domain.Profile) error
	Delete(ctx context.Context, userID int64) error
}

type profileRepository struct {
	conn *gorm.DB
}

func NewProfileRepository(conn *gorm.DB) ProfileRepository {
	return &profileRepository{conn}
}

func (r *profileRepository) GetByID(ctx context.Context, userID int64) (*domain.Profile, error) {
	var profile domain.Profile
	if err := r.conn.Where("user_id = ?", userID).Take(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *profileRepository) Create(ctx context.Context, p *domain.Profile) error {
	if err := r.conn.Create(&p).Error; err != nil {
		return err
	}
	return nil
}

func (r *profileRepository) Update(ctx context.Context, p *domain.Profile) error {
	var profile domain.Profile
	if err := r.conn.Where("user_id = ?", p.UserID).Take(&profile).Updates(p).Error; err != nil {
		return err
	}
	p.UpdatedAt = profile.UpdatedAt
	p.CreatedAt = profile.CreatedAt
	return nil
}

func (r *profileRepository) Delete(ctx context.Context, userID int64) error {
	var profile domain.Profile
	if err := r.conn.Where("user_id = ?", userID).Take(&profile).Delete(&profile).Error; err != nil {
		return err
	}
	return nil
}
