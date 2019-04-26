package model

import (
	"github.com/jinzhu/gorm"
)

type Model struct {
	DB *gorm.DB
}

func (M *Model) New(db *gorm.DB) {
	M.DB = Migrate(db)
}

func Migrate(db *gorm.DB)  *gorm.DB {
	db.AutoMigrate(&Category{})
	return db
}