package database

import "gorm.io/gorm"

var DB *gorm.DB

func Open() (err error) {
	DB, err = gorm.Open(noOpDialector{})
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
