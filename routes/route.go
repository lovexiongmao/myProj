package routes

import (
	"myProj/clients"
	_ "myProj/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title	myProj API
// @version	1.0
// @description	myProj后端api文档服务.
// @BasePath  /api/v1
// @schemes http https
// @securityDefinitions.basic BasicAuth
func SetUpRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	test := r.Group("api/test")
	test.GET("ping", clients.Ping)

	v1 := r.Group("api/v1")
	v1.POST("addUser", clients.AddUser)
	v1.DELETE("deleteUser", clients.DeleteUser)
	v1.PUT("updateUser", clients.UpDateUser)
	v1.GET("getUser", clients.GetUser)

	return r
}
