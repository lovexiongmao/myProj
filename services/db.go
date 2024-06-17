package services

import (
	log "myProj/sdk/log"
	"sync"

	conf "myProj/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func InitDB() *gorm.DB {
	once.Do(func() {
		config, err := conf.LoadConfig()
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}

		dsn := config.DB.DSN
		database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
		db = database
	})

	return db
}

func GetDB() *gorm.DB {
	return db
}
