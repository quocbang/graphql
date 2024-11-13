package repositories

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	Host         string
	Port         int
	UserName     string
	Password     string
	DatabaseName string
}

func NewDBConnection(cfg DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", cfg.Host, cfg.UserName, cfg.Password, cfg.DatabaseName, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
