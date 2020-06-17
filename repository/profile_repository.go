package repository

import (
	"context"

	"github.com/ezio1119/fishapp-profile/domain"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Usecase
type ProfileRepository interface {
	GetProfileByID(ctx context.Context, id int64) (*domain.Profile, error)
	GetProfileByUserID(ctx context.Context, userID int64) (*domain.Profile, error)
	BatchGetProfilesByUserIDs(ctx context.Context, userIDs []int64) ([]*domain.Profile, error)
	UpdateProfile(ctx context.Context, p *domain.Profile) error
	CreateProfile(ctx context.Context, p *domain.Profile) error
	DeleteProfile(ctx context.Context, userID int64) error
}

type profileRepository struct {
	conn *gorm.DB
}

func NewProfileRepository(conn *gorm.DB) ProfileRepository {
	return &profileRepository{conn}
}

func (r *profileRepository) GetProfileByID(ctx context.Context, id int64) (*domain.Profile, error) {
	p := &domain.Profile{ID: id}
	if err := r.conn.Take(p).Error; err != nil {
		return nil, err
	}

	return p, nil
}

func (r *profileRepository) GetProfileByUserID(ctx context.Context, uID int64) (*domain.Profile, error) {
	p := &domain.Profile{}
	if err := r.conn.Where("user_id = ?", uID).Take(p).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = status.Errorf(codes.NotFound, "profile with user_id='%d' is not found", uID)
		}
		return nil, err
	}
	return p, nil
}

func (r *profileRepository) BatchGetProfilesByUserIDs(ctx context.Context, userIDs []int64) ([]*domain.Profile, error) {
	p := []*domain.Profile{}
	if err := r.conn.Where("user_id IN (?)", userIDs).Find(&p).Error; err != nil {
		return nil, err
	}

	return p, nil
}

func (r *profileRepository) CreateProfile(ctx context.Context, p *domain.Profile) error {
	result := r.conn.Create(p)
	if err := result.Error; err != nil {
		e, ok := err.(*mysql.MySQLError)
		if ok {
			if e.Number == 1062 {
				err = status.Error(codes.AlreadyExists, err.Error())
			}
		}
		return err
	}

	return nil
}

func (r *profileRepository) UpdateProfile(ctx context.Context, p *domain.Profile) error {
	if err := r.conn.Model(p).Updates(p).Error; err != nil {
		return err
	}

	if err := r.conn.Take(p).Error; err != nil {
		return err
	}

	return nil
}

func (r *profileRepository) DeleteProfile(ctx context.Context, id int64) error {
	return r.conn.Delete(&domain.Profile{ID: id}).Error
}
