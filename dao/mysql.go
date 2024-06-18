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

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		// Add other models here
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
