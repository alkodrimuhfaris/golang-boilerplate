package services

import Models "papitupi-web/src/models"

type User struct {
}

type UserInterface interface {
	Login(us *Models.User) (*Models.LoggedIn, error)
}

func (u *User) Login(us *Models.User) (*Models.LoggedIn, error) {
	panic("implement me!")
}
