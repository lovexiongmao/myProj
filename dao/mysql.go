package dao

import (
	"myProj/models"
	log "myProj/sdk/log"

	"myProj/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(conf *conf.Config) *gorm.DB {
	dsn := conf.DB.DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func MigrateDB() {
	c, err := conf.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to get config %v", err)
	}
	db := NewDB(c)
	errMigrate := db.AutoMigrate(
		// Add other models here
		&models.User{},
		&models.AccountInfo{},
	)
	if errMigrate != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	defer func() {
		dbConnect, err := db.DB()
		if err != nil {
			log.Fatalf("Faild to open mysql connect: %v", err)
		}
		dbConnect.Close()
	}()
}
