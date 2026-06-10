package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Images      []ProductImage `json:"Images" gorm:"foreignKey:ProductID"`
}

type ProductImage struct {
	gorm.Model
	ProductID uint   `json:"product_id"`
	ImageURL  string `json:"image_url"`
}
