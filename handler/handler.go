package handler

import (
	"github.com/devlucas-java/api-opportunities/config"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *config.Logger
)

func InitializeHandler() {

	logger = config.GetLogger("handler")
	db = config.GetSQLiteDB()

}
