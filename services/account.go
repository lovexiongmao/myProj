package services

import (
	"errors"
	"fmt"
	"myProj/conf"
	"myProj/models"
	"myProj/sdk/log"
	"myProj/utils"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type AccountService struct {
	db   *gorm.DB
	conf *conf.Config
}

func NewAccountService(db *gorm.DB, conf *conf.Config) *AccountService {
	return &AccountService{
		db:   db,
		conf: conf,
	}
}

func (a *AccountService) AddAccount(u models.AccountInfo) error {
	if err := a.db.Create(&u).Error; err != nil {
		log.Errorf("create a account failed %v", err)
		return err
	}
	return nil
}

func (a *AccountService) Login(u models.AccountInfo) (map[string]string, error) {
	var user models.AccountInfo
	passwordCode := utils.GetBase64Code(u.Password)
	if err := a.db.Where("username = ? And password = ?", u.Username, passwordCode).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Errorf("Failed to login: %v", err)
			return nil, errors.New("用户名/密码错误！")
		}
		return nil, err
	}
	tokenStr, err := a.generateToken(u)
	if err != nil {
		return nil, err
	}
	token := map[string]string{
		"token": tokenStr,
	}
	return token, nil
}

func (a *AccountService) generateToken(u models.AccountInfo) (string, error) {
	accountClaim := &models.AccountClaims{
		Name:     u.Username,
		Identity: "identity",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, accountClaim)
	log.Info(a.conf.Secret)
	tokenStr, err := token.SignedString([]byte("nihao"))
	if err != nil {
		return "", fmt.Errorf("get token faild: %v", err.Error())
	}
	return tokenStr, nil

}
