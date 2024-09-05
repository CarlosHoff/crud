package config

import (
	"os"

	"github.com/CarlosHoff/crud.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	// Logger simulado (substitua com o seu)
	logger := GetLogger("sqlite")

	// Caminho do banco de dados SQLite
	dbPath := "./db/main.db"

	// Verifica se o arquivo de banco de dados existe
	_, err := os.Stat(dbPath)
	if err != nil && os.IsNotExist(err) {
		logger.Infof("db file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		defer file.Close()
	} else if err != nil {
		return nil, err
	}

	// Conecta ao banco de dados SQLite
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errf("sqlite opening error: %v", err)
		return nil, err
	}

	// Realiza a automigração dos esquemas
	err = db.AutoMigrate(&schemas.Crud{})
	if err != nil {
		logger.Errf("sqlite automigrate error: %v", err)
		return nil, err
	}

	return db, nil
}
