package services

import (
	"myProj/models"
	log "myProj/sdk/log"
)

func AddUser(user models.User) error {
	if err := GetDB().Create(&user).Error; err != nil {
		log.Fatalf("Failed to add user: %v", err)
		return err
	}
	return nil
}

func DeleteUser(name string) error {
	if err := GetDB().Where("name = ?", name).Delete(&models.User{}).Error; err != nil {
		log.Fatalf("Failed to delete user: %v", err)
		return err
	}
	return nil
}

func UpdateUser(user models.User) error {
	if err := GetDB().Save(&user).Error; err != nil {
		log.Fatalf("Failed to update user: %v", err)
		return err
	}
	return nil
}

func GetUser(name string) (models.User, error) {
	var user models.User
	if err := GetDB().Where("name = ?", name).First(&user).Error; err != nil {
		log.Fatalf("Failed to get user: %v", err)
		return user, err
	}
	return user, nil
}
