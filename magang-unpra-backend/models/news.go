package models

import "gorm.io/gorm"

type News struct {
	gorm.Model
	Title    string      `json:"title"`
	Content  string      `json:"content"`
	Category string      `json:"category"`
	Images   []NewsImage `json:"images"`
}

type NewsImage struct {
	gorm.Model
	NewsID   uint   `json:"news_id"`
	ImageURL string `json:"image_url"`
}
