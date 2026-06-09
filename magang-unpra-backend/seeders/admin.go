package seeders

import (
	"log"
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"

	"golang.org/x/crypto/bcrypt"
)

func SeedAdmin() {
	var count int64
	config.DB.Model(&models.User{}).Count(&count)
	if count > 0 {
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	admin := models.User{
		Name:     "Administrator",
		Email:    "admin@telpp.com",
		Password: string(hash),
		Role:     "admin",
	}

	config.DB.Create(&admin)
	log.Println("Admin seeded: admin@telpp.com / admin123")
}
