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
	var sex profile_grpc.Sex
	switch p.Sex {
	case domain.Male:
		sex = profile_grpc.Sex_MALE
	case domain.Female:
		sex = profile_grpc.Sex_FEMALE
	}
	return &profile_grpc.Profile{
		Id:           p.ID,
		Name:         p.Name,
		UserId:       p.UserID,
		Sex:          sex,
		Introduction: p.Introduction,
		UpdatedAt:    updatedAt,
		CreatedAt:    createdAt,
	}, nil
}

func convListProfilesProto(list []*domain.Profile) ([]*profile_grpc.Profile, error) {
	listP := make([]*profile_grpc.Profile, len(list))
	for i, p := range list {
		pProto, err := convProfileProto(p)
		if err != nil {
			return nil, err
		}
		listP[i] = pProto
	}
	return listP, nil
}
