package registry

import (
	"github.com/ezio1119/fishapp-profile/controllers"
	"github.com/ezio1119/fishapp-profile/interactor"
	"github.com/ezio1119/fishapp-profile/pb"
	"github.com/ezio1119/fishapp-profile/presenter"
	"github.com/ezio1119/fishapp-profile/repository"
)

func (r *registry) NewProfileController() pb.ProfileServiceServer {
	return controllers.NewProfileController(
		interactor.NewProfileInteractor(
			repository.NewProfileRepository(r.db),
			presenter.NewProfilePresenter(),
			r.timeout,
		))
}
