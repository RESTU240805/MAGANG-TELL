package handlers

import (
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
	"net/http"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
)

// helper buat slug tanpa library tambahan
func makeSlug(title string) string {
	var result strings.Builder
	for _, r := range strings.ToLower(title) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result.WriteRune(r)
		} else if unicode.IsSpace(r) {
			result.WriteRune('-')
		}
	}
	return result.String()
}

// Public - hanya yang published
func GetAllNews(c *gin.Context) {
	var news []models.News
	config.DB.Preload("Images").Where("is_published = ?", true).Find(&news)
	c.JSON(http.StatusOK, gin.H{"data": news})
}

// Admin - semua berita
func GetAllNewsAdmin(c *gin.Context) {
	var news []models.News
	config.DB.Preload("Images").Find(&news)
	c.JSON(http.StatusOK, gin.H{"data": news})
}

func GetNewsById(c *gin.Context) {
	var news models.News
	id := c.Param("id")
	if err := config.DB.Preload("Images").First(&news, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": news})
}

func CreateNews(c *gin.Context) {
	var news models.News
	if err := c.ShouldBindJSON(&news); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if news.Slug == "" {
		news.Slug = makeSlug(news.Title)
	}
	config.DB.Create(&news)
	c.JSON(http.StatusCreated, gin.H{"data": news})
}

func UpdateNews(c *gin.Context) {
	var news models.News
	id := c.Param("id")
	if err := config.DB.First(&news, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}
	c.ShouldBindJSON(&news)
	if news.Slug == "" {
		news.Slug = makeSlug(news.Title)
	}
	config.DB.Save(&news)
	c.JSON(http.StatusOK, gin.H{"data": news})
}

func DeleteNews(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.News{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "News deleted"})
}
