package restapi

import (
	"net/http"
	"papitupi-web/src/applications"
	"papitupi-web/src/models"
	"time"

	"github.com/gin-gonic/gin"
)

type RestAPIInfra struct {
	UserApplication applications.ApplicationsInterface
}

type RestAPIInfraInterface interface {
	Login(c *gin.Context)
}

func (u *RestAPIInfra) Login(c *gin.Context) {
	var password, username *string
	*password = c.Query("password")
	*username = c.Query("username")

	User := &models.User{
		Username: username,
		Password: password,
	}

	LoggedIn, err := u.UserApplication.Login(User)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
		return
	}

	maxAge := int(24 * time.Hour)
	domain := "localhost:3000"

	c.SetCookie(
		"access_token",
		LoggedIn.AccessToken,
		maxAge,
		"/",
		domain,
		true,
		true,
	)
	c.SetCookie(
		"refresh_token",
		LoggedIn.RefreshToken,
		maxAge,
		"/",
		domain,
		true,
		true,
	)
	c.SetCookie(
		"id_token",
		LoggedIn.IdToken,
		maxAge,
		"/",
		domain,
		true,
		true,
	)

	c.JSON(http.StatusOK, nil)
}
