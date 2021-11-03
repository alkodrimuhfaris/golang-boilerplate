package applications

import (
	"papitupi-web/src/models"
	"papitupi-web/src/services"
)

type UserApplications struct {
	UserServices services.UserInterface
}

type ApplicationsInterface interface {
	Login(u *models.User) (*models.LoggedIn, error)
}

func (a *UserApplications) Login(u *models.User) (*models.LoggedIn, error) {
	loggedInData, err := a.UserServices.Login(u)
	if err != nil {
		return nil, err
	}
	return loggedInData, nil
}
