package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return errors.New("fake Errors")
	//return nil
}

func GetLogger(v string) *Logger {
	logger = NewLogger(v)
	return logger
}
