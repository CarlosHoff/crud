package config

import (
	"os"

	"github.com/CarlosHoff/crud.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {

	logger := GetLogger("sqlite")
	dbPatch := "./db/main.db"

	_, err := os.Stat(dbPatch)
	if os.IsNotExist(err) {
		logger.Infof("DB file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPatch)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPatch), &gorm.Config{})
	if err != nil {
		logger.Errf("Sqlite opning error: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Crud{})
	if err != nil {
		logger.Errf("Sqlite automigrate error: %v", err)
		return nil, err
	}
	return db, nil
}
