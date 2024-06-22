package apiservice

import "github.com/gin-gonic/gin"

func (api *apiService) initDefaultRoute(v1 *gin.RouterGroup) {
	account := v1.Group("/account")
	account.POST("register", api.accountHandler.AddAccount)
	account.POST("login", api.accountHandler.Login)
}
