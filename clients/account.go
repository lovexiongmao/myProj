package clients

import (
	"errors"
	"myProj/common"
	"myProj/models"
	"myProj/sdk/log"
	"myProj/services"
	"myProj/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	account services.AccountService
}

func NewAccountHandler(accountService *services.AccountService) *AccountHandler {
	return &AccountHandler{
		account: *accountService,
	}
}

// addAccount interface
// @Summary add a account
// @Schemes
// @Description 注册一个平台用户
// @Tags account
// @param accountInfo body models.AccountInfo true "新增的用户信息"
// @Accept json
// @Produce json
// @Success 200
// @Router /account/register [post]
func (a *AccountHandler) AddAccount(ctx *gin.Context) {
	var accountInfo models.AccountInfo
	log.Debug(ctx.Params)
	if err := ctx.ShouldBindJSON(&accountInfo); err != nil {
		log.Errorf("please enter the correct content: %v", err)
		ctx.JSON(http.StatusBadRequest, common.ErrResp(err))
		return
	}
	passwordCode := utils.GetBase64Code(accountInfo.Password)
	accountInfo.Password = passwordCode
	if err := a.account.AddAccount(accountInfo); err != nil {
		ctx.JSON(http.StatusInternalServerError, common.ErrResp(err))
		return
	}
	ctx.JSON(http.StatusOK, "add account success! ")

}

// login interface
// @Summary login
// @Schemes
// @Description 登录平台
// @Tags account
// @param login body models.AccountInfo true "新增的用户信息"
// @Accept json
// @Produce json
// @Success 200
// @Router /account/login [post]
func (a *AccountHandler) Login(ctx *gin.Context) {
	var accountInfo models.AccountInfo
	log.Debug(ctx.Params)
	if err := ctx.ShouldBindJSON(&accountInfo); err != nil {
		log.Errorf("please enter the correct content: %v", err)
		ctx.JSON(http.StatusBadRequest, common.ErrResp(err))
		return
	}
	if err := a.account.Login(accountInfo); err != nil {
		err = errors.New("用户名/密码错误")
		ctx.JSON(http.StatusInternalServerError, common.ErrResp(err))
		return
	}
	ctx.JSON(http.StatusOK, "登录成功！")

}
