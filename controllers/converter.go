package controllers

import (
	"github.com/ezio1119/fishapp-profile/domain"
	"github.com/ezio1119/fishapp-profile/pb"
	"github.com/golang/protobuf/ptypes"
)

func convProfileProto(p *domain.Profile) (*pb.Profile, error) {
	updatedAt, err := ptypes.TimestampProto(p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	createdAt, err := ptypes.TimestampProto(p.CreatedAt)
	if err != nil {
		return nil, err
	}
	var sex pb.Sex
	switch p.Sex {
	case domain.Male:
		sex = pb.Sex_MALE
	case domain.Female:
		sex = pb.Sex_FEMALE
	}
	return &pb.Profile{
		Id:           p.ID,
		Name:         p.Name,
		UserId:       p.UserID,
		Sex:          sex,
		Introduction: p.Introduction,
		UpdatedAt:    updatedAt,
		CreatedAt:    createdAt,
	}, nil
}

func convListProfilesProto(list []*domain.Profile) ([]*pb.Profile, error) {
	listP := make([]*pb.Profile, len(list))
	for i, p := range list {
		pProto, err := convProfileProto(p)
		if err != nil {
			return nil, err
		}
		listP[i] = pProto
	}
	return listP, nil
}
