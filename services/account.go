package services

import (
	"myProj/models"
	"myProj/sdk/log"
	"myProj/utils"

	"gorm.io/gorm"
)

type AccountService struct {
	db *gorm.DB
}

func NewAccountService(db *gorm.DB) *AccountService {
	return &AccountService{
		db: db,
	}
}

func (a *AccountService) AddAccount(u models.AccountInfo) error {
	if err := a.db.Create(&u).Error; err != nil {
		log.Errorf("create a account failed %v", err)
		return err
	}
	return nil
}

func (a *AccountService) Login(u models.AccountInfo) error {
	var user models.AccountInfo
	passwordCode := utils.GetBase64Code(u.Password)
	if err := a.db.Where("username = ? And password = ?", u.Username, passwordCode).First(&user).Error; err != nil {
		log.Errorf("Failed to login: %v", err)
		return err
	}
	return nil
}
