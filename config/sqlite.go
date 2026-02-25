package config

import (
	"os"

	"github.com/devlucas-java/api-opportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db"
	dbFile := "/main.db"

	// Check if the datbase file no exists
	_, err := os.Stat(dbPath + dbFile)
	if os.IsNotExist(err) {

		logger.Info("File db not exists, creating file")
		// Create the directory if it doesn't exist
		err = os.MkdirAll(dbPath, os.ModePerm)
		if err != nil {
			logger.Errorf("Error creating db directory: %v", err)
			return nil, err
		}
		// Create the file
		file, err := os.Create(dbPath + dbFile)
		if err != nil {
			logger.Errorf("Error creating db file: %v", err)
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath+dbFile), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error connecting to SQLite database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Error migrating database: %v", err)
		return nil, err
	}

	return db, nil
}
