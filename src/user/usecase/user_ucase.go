package usecase

import (
	"context"
	"time"

	"github.com/ezio1119/fishapp-user/models"
	"github.com/ezio1119/fishapp-user/user"
)

type userUsecase struct {
	userRepo       user.Repository
	contextTimeout time.Duration
}

// NewuserUsecase will create new an userUsecase object representation of user.Usecase interface
func NewUserUsecase(p user.Repository, timeout time.Duration) user.Usecase {
	return &userUsecase{
		userRepo:       p,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) GetByID(ctx context.Context, id int64) (*models.User, error) {

	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	res, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *userUsecase) Create(ctx context.Context, user *models.User) (*models.TokenPair, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	encryptedPass, err := u.genEncryptedPass(user.Password)
	if err != nil {
		return nil, err
	}
	user.EncryptedPassword = encryptedPass
	if err := u.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	tokenPair, err := u.generateTokenPair(user.ID)
	if err != nil {
		return nil, err
	}
	return tokenPair, nil
}

func (u *userUsecase) Update(ctx context.Context, user *models.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	encryptedPass, err := u.genEncryptedPass(user.Password)
	if err != nil {
		return err
	}
	user.EncryptedPassword = encryptedPass
	if err := u.userRepo.Update(ctx, user); err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) Delete(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	return u.userRepo.Delete(ctx, id)
}

func (u *userUsecase) RefreshIDToken(ctx context.Context, rt string) (*models.TokenPair, error) {
	userID, err := u.validateToken(rt)
	tokenPair, err := u.generateTokenPair(userID)
	if err != nil {
		return nil, err
	}
	return tokenPair, nil
}

func (u *userUsecase) Login(ctx context.Context, email string, pass string) (*models.User, *models.TokenPair, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, nil, err
	}
	if err := u.compareHashAndPass(user.EncryptedPassword, pass); err != nil {
		return nil, nil, err
	}
	tokenPair, err := u.generateTokenPair(user.ID)
	if err != nil {
		return nil, nil, err
	}

	return user, tokenPair, nil
}
