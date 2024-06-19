package clients

import (
	"net/http"

	"errors"
	"myProj/common"
	"myProj/models"
	"myProj/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (user *UserHandler) Ping(ctx *gin.Context) {
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
// @Router /user/addUser [post]
func (u *UserHandler) AddUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrResp(err))
		return
	}
	if err := u.UserService.AddUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrResp(err))
		return
	}
	ctx.JSON(http.StatusOK, "add user")

}

// @Summary example for delete a user
// @Schemes
// @Description 测试的接口
// @Tags Test
// @param username query string  true "删除某个用户"
// @Accept json
// @Produce json
// @Success 200
// @Router /user/deleteUser [delete]
func (u *UserHandler) DeleteUser(ctx *gin.Context) {
	username := ctx.Query("username")
	if username == "" {
		ctx.JSON(http.StatusBadRequest, common.ErrResp(errors.New("请输入正确的参数！")))
		return
	}
	if err := u.UserService.DeleteUser(username); err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrResp(err))
		return
	}
	ctx.JSON(http.StatusOK, "delete user")
}

// @Summary example for update a user
// @Schemes
// @Description 测试的接口
// @Tags Test
// @param user body models.User true "更新的用户信息"
// @Accept json
// @Produce json
// @Success 200
// @Router /user/updateUser [put]
func (u *UserHandler) UpDateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrResp(err))
		return
	}
	if err := u.UserService.UpdateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrResp(err))
		return
	}
	ctx.JSON(http.StatusOK, "update user")
}

// @Summary example for get a user
// @Schemes
// @Description 测试的接口
// @Tags Test
// @param username query string  true "根据用户名来查询用户"
// @Accept json
// @Produce json
// @Success 200
// @Router /user/getUser [get]
func (u *UserHandler) GetUser(ctx *gin.Context) {
	username := ctx.Query("username")
	if username == "" {
		ctx.JSON(http.StatusBadRequest, common.ErrResp(errors.New("请输入正确的参数！")))
		return
	}
	user, err := u.UserService.GetUser(username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrResp(err))
		return
	}
	ctx.JSON(http.StatusOK, common.OkResp(user))
}
