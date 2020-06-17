package interactor

import (
	"context"
	"time"

	"github.com/ezio1119/fishapp-profile/domain"
	"github.com/ezio1119/fishapp-profile/presenter"
	"github.com/ezio1119/fishapp-profile/repository"
)

// Usecase
type ProfileInteractor interface {
	GetProfile(ctx context.Context, userID int64) (*domain.Profile, error)
	BatchGetProfiles(ctx context.Context, userIDs []int64) ([]*domain.Profile, error)
	UpdateProfile(ctx context.Context, p *domain.Profile) error
	CreateProfile(ctx context.Context, p *domain.Profile) error
	DeleteProfile(ctx context.Context, userID int64) error
}

type profileInteractor struct {
	profileRepository repository.ProfileRepository
	profilePresenter  presenter.ProfilePresenter
	ctxTimeout        time.Duration
}

func NewProfileInteractor(r repository.ProfileRepository, p presenter.ProfilePresenter, t time.Duration) ProfileInteractor {
	return &profileInteractor{r, p, t}
}

func (i *profileInteractor) GetProfile(ctx context.Context, userID int64) (*domain.Profile, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	return i.profileRepository.GetProfileByUserID(ctx, userID)
}

func (i *profileInteractor) BatchGetProfiles(ctx context.Context, userIDs []int64) ([]*domain.Profile, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	return i.profileRepository.BatchGetProfilesByUserIDs(ctx, userIDs)
}

func (i *profileInteractor) CreateProfile(ctx context.Context, p *domain.Profile) error {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	return i.profileRepository.CreateProfile(ctx, p)
}

func (i *profileInteractor) UpdateProfile(ctx context.Context, p *domain.Profile) error {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	return i.profileRepository.UpdateProfile(ctx, p)
}

func (i *profileInteractor) DeleteProfile(ctx context.Context, uID int64) error {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	res, err := i.profileRepository.GetProfileByUserID(ctx, uID)
	if err != nil {
		return err
	}
	return i.profileRepository.DeleteProfile(ctx, res.ID)
}
