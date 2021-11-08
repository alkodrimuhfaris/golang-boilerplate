package keycloakService

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"papitupi-web/src/models"
	"strconv"
	"strings"

	"github.com/ajg/form"
)

type KeycloakService struct {
	Cl              *http.Client
	Tr              *http.Transport
	KeycloakBaseURL string
	ClientID        string
	ClientSecret    string
}

func (k *KeycloakService) Login(us *models.Credential) (*models.LoggedIn, error) {
	log.Println("login with keycloak")

	keycloakPayload := Credential{
		Password:     *us.Password,
		Username:     *us.Username,
		ClientID:     k.ClientID,
		ClientSecret: k.ClientSecret,
		Scope:        "openid",
		GrantType:    "password",
	}

	k.Tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	k.Cl.Transport = k.Tr

	payloadURLValues, err := form.EncodeToValues(keycloakPayload)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	payloadIOString := strings.NewReader(payloadURLValues.Encode())

	r, err := http.NewRequest("POST", k.KeycloakBaseURL, payloadIOString)
	if err != nil {
		log.Println("Error on creating post request to keycloak")
		log.Println(err)
		return nil, err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(payloadURLValues.Encode())))

	res, err := k.Cl.Do(r)
	if err != nil {
		log.Println("Error on doing post request to keycloak")
		log.Println(err)
		return nil, err
	}
	log.Println(res.Status)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error on read body response from keycloak login")
		log.Println(err)
		return nil, err
	}

	resp := &Response{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Println("Error marshalling json response from keycloak login")
		log.Println(err)
		return nil, err
	}

	loggedIn := &models.LoggedIn{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
		IdToken:      resp.IDToken,
	}
	return loggedIn, nil
}
