package main

import (
	"kilkenny/purpleschool/configs"
	"kilkenny/purpleschool/internal/link"
	"kilkenny/purpleschool/pkg/db"
)

func main() {
	config := configs.LoadConfig()
	db := db.NewDb(config)

	db.AutoMigrate(&link.Link{})
}
