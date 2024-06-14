package routes

import (
	"myProj/clients"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) *gin.Engine {
	r.Group("api/v1")
	r.GET("test", clients.MyTest)
	return r
}
