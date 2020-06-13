package controllers

import (
	"context"

	"github.com/ezio1119/fishapp-profile/domain"
	"github.com/ezio1119/fishapp-profile/interactor"
	"github.com/ezio1119/fishapp-profile/pb"
	"github.com/golang/protobuf/ptypes/empty"
)

type profileController struct {
	profileInteractor interactor.ProfileInteractor
}

func NewProfileController(i interactor.ProfileInteractor) pb.ProfileServiceServer {
	return &profileController{i}
}

func (c *profileController) CreateProfile(ctx context.Context, in *pb.CreateProfileReq) (*pb.Profile, error) {
	p := &domain.Profile{
		Name:         in.Name,
		Introduction: in.Introduction,
		UserID:       in.UserId,
	}
	switch in.Sex {
	case pb.Sex_MALE:
		p.Sex = domain.Male
	case pb.Sex_FEMALE:
		p.Sex = domain.Female
	}
	if err := c.profileInteractor.CreateProfile(ctx, p); err != nil {
		return nil, err
	}
	return convProfileProto(p)
}

func (c *profileController) GetProfile(ctx context.Context, in *pb.GetProfileReq) (*pb.Profile, error) {
	p, err := c.profileInteractor.GetProfile(ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	return convProfileProto(p)
}

func (c *profileController) BatchGetProfiles(ctx context.Context, in *pb.BatchGetProfilesReq) (*pb.BatchGetProfilesRes, error) {
	p, err := c.profileInteractor.BatchGetProfiles(ctx, in.UserIds)
	if err != nil {
		return nil, err
	}
	pProto, err := convListProfilesProto(p)
	if err != nil {
		return nil, err
	}
	return &pb.BatchGetProfilesRes{
		Profiles: pProto,
	}, nil
}

func (c *profileController) UpdateProfile(ctx context.Context, in *pb.UpdateProfileReq) (*pb.Profile, error) {
	p := &domain.Profile{
		Name:         in.Name,
		Introduction: in.Introduction,
		UserID:       in.UserId,
	}
	if err := c.profileInteractor.UpdateProfile(ctx, p); err != nil {
		return nil, err
	}
	return convProfileProto(p)
}

func (c *profileController) DeleteProfile(ctx context.Context, in *pb.DeleteProfileReq) (*empty.Empty, error) {
	if err := c.profileInteractor.DeleteProfile(ctx, in.UserId); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
