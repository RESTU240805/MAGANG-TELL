package seeders

import (
	"log"
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
)

func SeedAdmin() {
	var count int64
	config.DB.Model(&models.User{}).Count(&count)
	if count > 0 {
		return
	}

	admin := models.User{
		Name:     "Administrator",
		Email:    "admin@telpp.com",
		Password: "admin123",
		Role:     "admin",
	}

	config.DB.Create(&admin)
	log.Println("Admin seeded: admin@telpp.com / admin123")
}
