package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adrianosiqe/eulabs-challenge-api/internal/domains/models"
	"github.com/adrianosiqe/eulabs-challenge-api/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		config.Cfg.DBUser, config.Cfg.DBPassword, config.Cfg.DBHost, config.Cfg.DBPort, config.Cfg.DBName, config.Cfg.DBCharset, config.Cfg.DBParseTime, config.Cfg.DBLoc)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Error,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Product{})

	return db, nil
}
