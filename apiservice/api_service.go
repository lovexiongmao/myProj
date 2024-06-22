package apiservice

import (
	"myProj/clients"

	_ "myProj/docs"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type ApiServiceInterface interface {
	Run(addr string)
}

type params struct {
	dig.In
	UserHandler    *clients.UserHandler
	AccountHandler *clients.AccountHandler
}

type apiService struct {
	route          *gin.Engine
	userHandler    *clients.UserHandler
	accountHandler *clients.AccountHandler
}

func NewApiService(p params) ApiServiceInterface {
	g := gin.Default()
	api := &apiService{
		route:          g,
		userHandler:    p.UserHandler,
		accountHandler: p.AccountHandler,
	}
	api.initRoute()
	return api
}

func (api *apiService) Run(addr string) {
	api.route.Run(addr)
}
