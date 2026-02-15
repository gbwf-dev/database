package database

import (
	"gbfw/api/env"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Open() (err error) {
	DB, err = gorm.Open(sqlite.Open(env.Getenv("DB_PATH", "database.sqlite")))
	return
}

func AutoMigrate() (err error) {
	return DB.AutoMigrate(
	// Insert auto migrated models here
	)
}

func Close() (err error) {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
