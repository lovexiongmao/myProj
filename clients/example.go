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
// @param name body string true "用户名" default(张三)
// @param age body int true "年龄" default(18)
// @param gender body string true "性别" default(男)
// @param hight body int true "身高" default(180)
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
// @Accept json
// @Produce json
// @Success 200
// @Router /getUser [get]
func GetUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "get user")
}
