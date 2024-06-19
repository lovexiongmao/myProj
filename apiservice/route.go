package apiservice

import (
	_ "myProj/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title	myProj API
// @version	1.0
// @description	myProj后端api文档服务.
// @BasePath  /api/v1
// @schemes http https
// @securityDefinitions.basic BasicAuth
func (api *apiService) initRoute() {
	r := api.route
	v1 := r.Group("api/v1")
	// swagger路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 测试路由
	test := r.Group("api/test")
	test.GET("ping", api.userHandler.Ping)

	// 各个业务的请求路由
	api.initDefaultRoute(v1)
	api.initExampleRoute(v1)

}
