/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2024-06-14 23:20:06
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2024-06-15 00:48:45
 * @FilePath: \myProj\clients\example.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package clients

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary ping example1
// @Schemes
// @Description 测试的接口1
// @Tags uTest
// @Accept json
// @Produce json
// @Success 200 {string} MyTest
// @Router /ping [put]
func MyTest(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ok")
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
}
