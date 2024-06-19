package apiservice

import "github.com/gin-gonic/gin"

func (api *apiService) initExampleRoute(v1 *gin.RouterGroup) {
	user := v1.Group("/user")
	user.POST("addUser", api.userHandler.AddUser)
	user.DELETE("deleteUser", api.userHandler.DeleteUser)
	user.PUT("updateUser", api.userHandler.UpDateUser)
	user.GET("getUser", api.userHandler.GetUser)
}
