package gorm

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"myapp/config"
)

func New(conf *config.DbConf) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		conf.Host,
		conf.Username,
		conf.Password,
		conf.DbName,
		conf.Port,
	)

	var logLevel logger.LogLevel
	if conf.Debug {
		logLevel = logger.Info
	} else {
		logLevel = logger.Error
	}

	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
}
