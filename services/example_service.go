package services

import (
	"myProj/models"
	log "myProj/sdk/log"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s UserService) AddUser(user models.User) error {
	if err := s.db.Create(&user).Error; err != nil {
		log.Errorf("Failed to add user: %v", err)
		return err
	}
	return nil
}

func (s UserService) DeleteUser(name string) error {
	if err := s.db.Where("name = ?", name).Delete(&models.User{}).Error; err != nil {
		log.Errorf("Failed to delete user: %v", err)
		return err
	}
	return nil
}

func (s UserService) UpdateUser(user models.User) error {
	if err := s.db.Model(&models.User{}).Where("name = ?", user.Name).Updates(user).Error; err != nil {
		log.Errorf("Failed to update user: %v", err)
		return err
	}
	return nil
}

func (s UserService) GetUser(name string) (models.User, error) {
	var user models.User
	if err := s.db.Where("name = ?", name).First(&user).Error; err != nil {
		log.Errorf("Failed to get user: %v", err)
		return user, err
	}
	return user, nil
}
