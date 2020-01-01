package presenter

import (
	"github.com/ezio1119/fishapp-profile/controllers/profile_grpc"
	"github.com/ezio1119/fishapp-profile/domain"
	"github.com/golang/protobuf/ptypes"
)

// Usecase
type ProfilePresenter interface {
	TransformProfileProto(*domain.Profile) (*profile_grpc.Profile, error)
}

type profilePresenter struct{}

func NewProfilePresenter() ProfilePresenter {
	return &profilePresenter{}
}

func (*profilePresenter) TransformProfileProto(p *domain.Profile) (*profile_grpc.Profile, error) {
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
