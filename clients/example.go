/*
 * @Author: lovexiongmao 953191418@qq.com
 * @Date: 2024-06-14 23:20:06
 * @LastEditors: lovexiongmao 953191418@qq.com
 * @LastEditTime: 2024-06-18 00:05:22
 * @FilePath: \myProj\clients\example.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package clients

import (
	"net/http"

	"myProj/common"
	"myProj/models"
	"myProj/services"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong")
}

// addUser interface
// @Summary example for add a user
// @Schemes
// @Description 测试的接口
// @Tags Test
// @param user body models.User true "新增的用户信息"
// @Accept json
// @Produce json
// @Success 200
// @Router /addUser [post]
func AddUser(ctx *gin.Context) {
	var req models.User
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, common.ErrResp(err))
		return
	}
	services.AddUser(req)
	ctx.JSON(http.StatusOK, "add user")
}

// @Summary example for delete a user
// @Schemes
// @Description 测试的接口
// @Tags Test
// @param name body string true "用户名" default(张三)
// @Accept json
// @Produce json
// @Success 200
// @Router /deleteUser [delete]
func DeleteUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "delete user")
}

// @Summary example for update a user
// @Schemes
// @Description 测试的接口
// @Tags Test
// @Accept json
// @Produce json
// @Success 200
// @Router /updateUser [put]
func UpDateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "update user")
}

// @Summary example for get a user
// @Schemes
// @Description 测试的接口
// @Tags Test
// @param user query string  true "根據用戶名來查用戶"
// @Accept json
// @Produce json
// @Success 200
// @Router /getUser [get]
func GetUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "get user")
}
