package restapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Router  *gin.Engine
	RestAPI RestAPIInfraInterface
}

type RouterInterface interface {
	Routes()
}

func (r *Router) Routes() {

	r.Router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "papitupi web service is running")
	})

	r.Router.POST("/login", r.RestAPI.Login)
}
