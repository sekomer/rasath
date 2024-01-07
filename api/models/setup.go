package models

import "gorm.io/gorm"

func Setup(db *gorm.DB) {
	db.Migrator().DropTable(
		&Earthquake{},
	)

	db.AutoMigrate(
		&Earthquake{},
	)
}
