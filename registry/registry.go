package registry

import (
	"time"

	"github.com/ezio1119/fishapp-profile/pb"
	"github.com/jinzhu/gorm"
)

type registry struct {
	db      *gorm.DB
	timeout time.Duration
}

type Registry interface {
	NewProfileController() pb.ProfileServiceServer
}

func NewRegistry(db *gorm.DB, t time.Duration) Registry {
	return &registry{db, t}
}
