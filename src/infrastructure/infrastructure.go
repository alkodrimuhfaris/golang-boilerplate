package infrastructure

import (
	"fmt"
	"net/http"
	"papitupi-web/src/applications"
	"papitupi-web/src/infrastructure/keycloakService"
	"papitupi-web/src/infrastructure/restapi"
	"papitupi-web/src/services"
	"time"

	"github.com/gin-gonic/gin"
)

func NewKeycloakService(Cl *http.Client) services.AuthServiceInterface {
	return &keycloakService.KeycloakService{Cl: Cl}
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
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	cl := &http.Client{Transport: tr}
	keycloak := NewKeycloakService(cl)
	applicationLayer := NewAplications(keycloak)
	restAPI := NewRestapiGinInfra(applicationLayer)

	// initialize gin engine
	router := gin.Default()
	api := NewRouteService(router, restAPI)
	api.Routes()
	router.Run(fmt.Sprintf(":%d", 7992))
}
