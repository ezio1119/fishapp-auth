package controllers

import (
	"context"

	"github.com/ezio1119/fishapp-profile/controllers/profile_grpc"
	"github.com/ezio1119/fishapp-profile/domain"
	"github.com/ezio1119/fishapp-profile/interactor"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type profileController struct {
	profileInteractor interactor.ProfileInteractor
}

func NewProfileController(i interactor.ProfileInteractor) profile_grpc.ProfileServiceServer {
	return &profileController{i}
}

func (c *profileController) Create(ctx context.Context, in *profile_grpc.CreateReq) (*profile_grpc.Profile, error) {
	profile := &domain.Profile{
		Name:   in.Name,
		UserID: in.UserId,
	}
	return c.profileInteractor.Create(ctx, profile)
}

func (c *profileController) GetByUserID(ctx context.Context, in *profile_grpc.ID) (*profile_grpc.Profile, error) {
	return c.profileInteractor.GetByUserID(ctx, in.UserId)
}

func (c *profileController) UpdateByUserID(ctx context.Context, in *profile_grpc.UpdateReq) (*profile_grpc.Profile, error) {
	profile := &domain.Profile{
		Name:   in.Name,
		UserID: in.UserId,
	}
	return c.profileInteractor.UpdateByUserID(ctx, profile)
}

func (c *profileController) DeleteByUserID(ctx context.Context, in *profile_grpc.ID) (*wrappers.BoolValue, error) {
	if err := c.profileInteractor.DeleteByUserID(ctx, in.UserId); err != nil {
		return &wrappers.BoolValue{Value: false}, err
	}
	return &wrappers.BoolValue{Value: true}, nil
}
