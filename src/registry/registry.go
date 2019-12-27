package registry

import (
	"time"

	"github.com/ezio1119/fishapp-auth/interfaces/controllers"
	"github.com/ezio1119/fishapp-auth/interfaces/repository"
	"github.com/ezio1119/fishapp-auth/usecase/interactor"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
)

func NewUserController(t time.Duration, conn *gorm.DB, client *redis.Client) *controllers.UserController {
	return &controllers.UserController{
		UserInteractor: &interactor.UserInteractor{
			UserRepository: &repository.UserRepository{
				Conn: conn,
			},
			BlackListRepository: &repository.BlackListRepository{
				Client: client,
			},
			TokenInteractor: &interactor.TokenInteractor{},
			ContextTimeout:  t,
		},
	}
}
