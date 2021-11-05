package applications

import (
	"papitupi-web/src/models"
	"papitupi-web/src/services"
)

type AuthApplications struct {
	AuthService services.AuthServiceInterface
}

type AuthApplicationsInterface interface {
	Login(u *models.Credential) (*models.LoggedIn, error)
}

func (a *AuthApplications) Login(u *models.Credential) (*models.LoggedIn, error) {
	loggedInData, err := a.AuthService.Login(u)
	if err != nil {
		return nil, err
	}
	return loggedInData, nil
}
