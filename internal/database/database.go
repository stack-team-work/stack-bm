package database

import (
	"fmt"
	"log"

	"stack-bm/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DBBM  *gorm.DB
	DBApi *gorm.DB
)

func InitDB() {
	var err error

	logLevel := logger.Info
	if config.AppConfig.Server.Mode == "prod" {
		logLevel = logger.Error
	}

	DBBM, err = gorm.Open(mysql.Open(config.AppConfig.DBBM.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database %s: %v", config.AppConfig.DBBM.Name, err)
	}
	fmt.Printf("Connected to database: %s\n", config.AppConfig.DBBM.Name)

	DBApi, err = gorm.Open(mysql.Open(config.AppConfig.DBApi.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database %s: %v", config.AppConfig.DBApi.Name, err)
	}
	fmt.Printf("Connected to database: %s\n", config.AppConfig.DBApi.Name)
}
