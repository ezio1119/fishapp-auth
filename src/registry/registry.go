package registry

import (
	"time"

	"github.com/ezio1119/fishapp-profile/controllers/profile_grpc"
	"github.com/jinzhu/gorm"
)

type registry struct {
	db      *gorm.DB
	timeout time.Duration
}

type Registry interface {
	NewProfileController() profile_grpc.ProfileServiceServer
}

func NewRegistry(db *gorm.DB, t time.Duration) Registry {
	return &registry{db, t}
}
