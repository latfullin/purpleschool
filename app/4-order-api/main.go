package main

import (
	"kilkenny/purpleschool/configs"
	"kilkenny/purpleschool/pkg/db"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	Name        string         `json:"name" gorm:"required"  validate:"required"`
	Description string         `json:"description" gorm:"required"  validate:"required"`
	Images      pq.StringArray `gorm:"type:text[]"`
}

func main() {
	config := configs.LoadConfig()

	db := db.NewDb(config)
	db.AutoMigrate(&Product{})
}
