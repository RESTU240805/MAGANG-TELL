package routes

import (
	"time"

	"magang-unpra-backend/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Serve uploaded files
	r.Static("/uploads", "./uploads")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/api/login", handlers.Login)
	r.POST("/api/upload", handlers.UploadImage)

	api := r.Group("/api")
	{
		// Public
		api.GET("/news", handlers.GetAllNews)
		api.GET("/news/:id", handlers.GetNewsById)
		api.GET("/products", handlers.GetAllProducts)
		api.GET("/products/:id", handlers.GetProductById)
		api.GET("/company-profile", handlers.GetCompanyProfile)

		// Admin
		admin := api.Group("/admin")
		{
			admin.GET("/news", handlers.GetAllNewsAdmin)
			admin.POST("/news", handlers.CreateNews)
			admin.PUT("/news/:id", handlers.UpdateNews)
			admin.DELETE("/news/:id", handlers.DeleteNews)

			admin.GET("/products", handlers.GetAllProducts)
			admin.POST("/products", handlers.CreateProduct)
			admin.PUT("/products/:id", handlers.UpdateProduct)
			admin.DELETE("/products/:id", handlers.DeleteProduct)

			admin.PUT("/company-profile", handlers.UpdateCompanyProfile)
		}
	}

	return r
}
