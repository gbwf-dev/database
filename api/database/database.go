package database

import (
	"fmt"
	"gbfw/api/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Open() (err error) {
	DB, err = gorm.Open(postgres.Open(fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		env.Getenv("DB_USER", "potgres"),
		env.Getenv("DB_PASSWORD", ""),
		env.Getenv("DB_HOST", "127.0.0.1"),
		env.Getenv("DB_PORT", "5432"),
		env.Getenv("DB_NAME", "gbwf"),
	)))
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
