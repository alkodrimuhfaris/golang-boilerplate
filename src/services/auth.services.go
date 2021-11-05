package services

import Models "papitupi-web/src/models"

type AuthService struct {
}

type AuthServiceInterface interface {
	Login(us *Models.Credential) (*Models.LoggedIn, error)
}

func (u *AuthService) Login(us *Models.Credential) (*Models.LoggedIn, error) {
	panic("implement me!")
}
