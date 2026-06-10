package handlers

import (
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllMenus berfungsi untuk mengambil semua data menu dari database
func GetAllMenus(c *gin.Context) {
	var menus []models.Menu

	// Mengambil semua data dari tabel menus menggunakan GORM
	if err := config.DB.Find(&menus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Mengembalikan data menu dalam format JSON dengan status 200 OK
	c.JSON(http.StatusOK, menus)
}
