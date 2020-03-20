package controllers

import (
	"github.com/ezio1119/fishapp-profile/controllers/profile_grpc"
	"github.com/ezio1119/fishapp-profile/domain"
	"github.com/golang/protobuf/ptypes"
)

func convProfileProto(p *domain.Profile) (*profile_grpc.Profile, error) {
	updatedAt, err := ptypes.TimestampProto(p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	createdAt, err := ptypes.TimestampProto(p.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &profile_grpc.Profile{
		Id:        p.ID,
		Name:      p.Name,
		UserId:    p.UserID,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}, nil
}
