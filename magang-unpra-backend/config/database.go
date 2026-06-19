package config

import (
	"fmt"
	"log"
	"os"

	"magang-unpra-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(
		&models.User{},
		&models.News{},
		&models.NewsImage{},
		&models.Product{},
		&models.ProductImage{},
		&models.ProductSlider{},

		&models.Creed{},
		&models.Menu{},
		&models.CompanyProfile{},
		&models.Sustainability{},
		&models.SustainabilityImage{},
		&models.ProductPage{},
		&models.AboutSection{},
		&models.TeamMember{},
	)

	db.AutoMigrate(
		&models.User{},
		&models.News{},
		&models.NewsImage{},
		&models.Product{},
		&models.ProductImage{},
		&models.ProductSlider{},
		&models.CompanyProfile{},
	)

	db.AutoMigrate(
		&models.User{},
		&models.News{},
		&models.NewsImage{},
		&models.Product{},
		&models.ProductImage{},
		&models.ProductSlider{},
		&models.Creed{},
		&models.Menu{},
		&models.CompanyProfile{},
		&models.CompanyDocument{}, // <-- tambahkan ini
		&models.Sustainability{},
		&models.SustainabilityImage{},
		&models.ProductPage{},
		&models.AboutSection{},
		&models.TeamMember{},
	)

	log.Println("Database connected successfully!")
	DB = db
}
