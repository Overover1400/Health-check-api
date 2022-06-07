package store

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySqlStore struct {
	db *gorm.DB
}

func New(dataBase string) MySqlStore {

	db, err := gorm.Open(mysql.Open(dataBase))
	if err != nil {
		panic("Failed to connect Database" + dataBase)
	}

	if err := db.AutoMigrate(&ClientApi{}); err != nil {
		panic("Failed to Auto Migrate database")
	}

	return MySqlStore{db}
}
