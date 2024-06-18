package apiservice

import (
	"myProj/clients"

	_ "myProj/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/dig"
)

type ApiServiceInterface interface {
	Run(addr string)
}

type params struct {
	dig.In
	UserHandler *clients.UserHandler
}

type apiService struct {
	route       *gin.Engine
	userHandler *clients.UserHandler
}

func NewApiService(p params) ApiServiceInterface {
	g := gin.Default()
	api := &apiService{
		route:       g,
		userHandler: p.UserHandler,
	}
	api.initRoute()
	return api
}

func (api *apiService) Run(addr string) {
	api.route.Run(addr)
}

// @title	myProj API
// @version	1.0
// @description	myProj后端api文档服务.
// @BasePath  /api/v1
// @schemes http https
// @securityDefinitions.basic BasicAuth
func (api *apiService) initRoute() {
	r := api.route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	test := r.Group("api/test")
	test.GET("ping", api.userHandler.Ping)

	v1 := r.Group("api/v1")
	v1.POST("addUser", api.userHandler.AddUser)
	v1.DELETE("deleteUser", api.userHandler.DeleteUser)
	v1.PUT("updateUser", api.userHandler.UpDateUser)
	v1.GET("getUser", api.userHandler.GetUser)
}
