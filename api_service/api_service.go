package apiservice

import (
	"myProj/routes"

	"github.com/gin-gonic/gin"
)

type ApiService interface {
	Run(string)
}

type Api struct {
	G *gin.Engine
}

func NewApiService() ApiService {
	r := gin.Default()
	route := routes.SetUpRoutes(r)
	api := Api{
		G: route,
	}
	return api
}

func (api Api) Run(port string) {
	api.G.Run(port)
}
