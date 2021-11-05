package keycloakService

import (
	"fmt"
	"net/http"
	"papitupi-web/src/models"
)

type KeycloakService struct {
	Cl *http.Client
}

func (k *KeycloakService) Login(us *models.Credential) (*models.LoggedIn, error) {
	fmt.Println("mantap slur")
	loggedIn := &models.LoggedIn{
		AccessToken:  "12314151523",
		RefreshToken: "12314151523",
		IdToken:      "12314151523",
	}
	return loggedIn, nil
}
