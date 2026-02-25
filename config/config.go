package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(v string) *Logger {
	logger = NewLogger(v)
	return logger
}
