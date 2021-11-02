package models

type User struct {
	Username string
	Password string
	Role     string
}

type LoggedIn struct {
	AccessToken  string
	RefreshToken string
	IdToken      string
}
