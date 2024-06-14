package clients

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MyTest(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ok")
	return
}

// @Summary ping example
// @Schemes
// @Description 测试的接口
// @Tags myTest
// @Accept json
// @Produce json
// @Success 200 {string} MyTest
// @Router /ping [get]
func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong")
	return
}
