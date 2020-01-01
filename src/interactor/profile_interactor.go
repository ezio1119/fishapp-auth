package interactor

import (
	"context"
	"time"

	"github.com/ezio1119/fishapp-profile/controllers/profile_grpc"
	"github.com/ezio1119/fishapp-profile/domain"
	"github.com/ezio1119/fishapp-profile/presenter"
	"github.com/ezio1119/fishapp-profile/repository"
)

// Usecase
type ProfileInteractor interface {
	GetByID(ctx context.Context, userID int64) (*profile_grpc.Profile, error)
	Update(ctx context.Context, p *domain.Profile) (*profile_grpc.Profile, error)
	Create(ctx context.Context, p *domain.Profile) (*profile_grpc.Profile, error)
	Delete(ctx context.Context, userID int64) error
}

type profileInteractor struct {
	profileRepository repository.ProfileRepository
	profilePresenter  presenter.ProfilePresenter
	ctxTimeout        time.Duration
}

func NewProfileInteractor(r repository.ProfileRepository, p presenter.ProfilePresenter, t time.Duration) ProfileInteractor {
	return &profileInteractor{r, p, t}
}

func (i *profileInteractor) GetByID(ctx context.Context, userID int64) (*profile_grpc.Profile, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	profile, err := i.profileRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return i.profilePresenter.TransformProfileProto(profile)
}

func (i *profileInteractor) Create(ctx context.Context, profile *domain.Profile) (*profile_grpc.Profile, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	if err := i.profileRepository.Create(ctx, profile); err != nil {
		return nil, err
	}
	return i.profilePresenter.TransformProfileProto(profile)
}

func (i *profileInteractor) Update(ctx context.Context, profile *domain.Profile) (*profile_grpc.Profile, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	if err := i.profileRepository.Update(ctx, profile); err != nil {
		return nil, err
	}
	return i.profilePresenter.TransformProfileProto(profile)
}

func (i *profileInteractor) Delete(ctx context.Context, userID int64) error {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	return i.profileRepository.Delete(ctx, userID)
}
