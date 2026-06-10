package handlers

import (
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"

	"github.com/gin-gonic/gin"
)

func GetAllCreeds(c *gin.Context) {
	var creeds []models.Creed

	if err := config.DB.Find(&creeds).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, creeds)
}
