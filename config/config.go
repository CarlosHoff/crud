package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error

	// Initialize DB Sqlite
	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("erro ao inicializar o sqlite: %w", err)
	}

	return nil
}

func GetSqlite() *gorm.DB {
	return db

}

func GetLogger(p string) *Logger {
	return NewLogger(p)
}
