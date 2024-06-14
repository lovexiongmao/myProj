package routes

import (
	"myProj/clients"

	docs "myProj/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpRoutes(r *gin.Engine) *gin.Engine {
	docs.SwaggerInfo.Title = "myProj API"
	docs.SwaggerInfo.Description = "myProj后端api文档服务"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	test := r.Group("api/test")
	test.GET("tmptest", clients.MyTest)

	v1 := r.Group("api/v1")
	v1.GET("ping", clients.Ping)

	return r
}
