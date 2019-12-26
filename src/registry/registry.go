package registry

import (
	"time"

	"github.com/ezio1119/fishapp-auth/interfaces/controllers"
	"github.com/ezio1119/fishapp-auth/interfaces/repository"
	"github.com/ezio1119/fishapp-auth/usecase/interactor"
	"github.com/jinzhu/gorm"
)

func NewUserController(t time.Duration, conn *gorm.DB) *controllers.UserController {
	return &controllers.UserController{
		UserInteractor: &interactor.UserInteractor{
			UserRepository: &repository.UserRepository{
				Conn: conn,
			},
			TokenInteractor: &interactor.TokenInteractor{},
			ContextTimeout:  t,
		},
	}
}
