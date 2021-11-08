package infrastructure

import (
	"fmt"
	"net/http"
	"papitupi-web/src/applications"
	"papitupi-web/src/infrastructure/keycloakService"
	"papitupi-web/src/infrastructure/restapi"
	conf "papitupi-web/src/infrastructure/viperConfig"
	"papitupi-web/src/services"
	"time"

	"github.com/gin-gonic/gin"
)

func NewSetConfig() conf.ConfigInterface {
	return &conf.Config{}
}

func NewKeycloakService(
	Cl *http.Client,
	Tr *http.Transport,
	KeycloakBaseURL string,
) services.AuthServiceInterface {
	return &keycloakService.KeycloakService{
		Cl:              Cl,
		Tr:              Tr,
		KeycloakBaseURL: KeycloakBaseURL,
	}
}

func NewAplications(
	AuthService services.AuthServiceInterface,
) applications.AuthApplicationsInterface {
	return &applications.AuthApplications{
		AuthService: AuthService,
	}
}

func NewRestapiGinInfra(
	AuthApplication applications.AuthApplicationsInterface,
) restapi.RestAPIInfraInterface {
	return &restapi.RestAPIInfra{
		AuthApplication: AuthApplication,
	}
}

func NewRouteService(
	Router *gin.Engine,
	RestAPI restapi.RestAPIInfraInterface,
) restapi.RouterInterface {
	return &restapi.Router{
		Router:  Router,
		RestAPI: RestAPI,
	}
}

func Infrastructure() {
	// initialize config
	newConf := NewSetConfig()
	config := newConf.SetConfig()

	// initialize http transport
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	// initialzie http client
	cl := &http.Client{Transport: tr}

	keycloak := NewKeycloakService(cl, tr, config.KeycloakBaseURL)
	applicationLayer := NewAplications(keycloak)
	restAPI := NewRestapiGinInfra(applicationLayer)

	// initialize gin engine
	router := gin.Default()
	api := NewRouteService(router, restAPI)
	api.Routes()
	router.Run(fmt.Sprintf(":%s", config.Port))
}
