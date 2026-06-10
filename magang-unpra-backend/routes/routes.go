package routes

import (
	"magang-unpra-backend/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// ✅ Tambah ini — serve folder uploads sebagai file statis
	r.Static("/uploads", "./uploads")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/api/login", handlers.Login)

	api := r.Group("/api")
	{
		// Public routes
		api.GET("/news", handlers.GetAllNews)
		api.GET("/news/:id", handlers.GetNewsById)

		api.GET("/products", handlers.GetAllProducts)
		api.GET("/products/:id", handlers.GetProductById)

		api.GET("/company-profile", handlers.GetCompanyProfile)

		// Admin routes
		admin := api.Group("/admin")
		{
			// ✅ Tambah ini — route upload gambar
			admin.POST("/upload/image", handlers.UploadProductImage)

			// News
			admin.GET("/news", handlers.GetAllNewsAdmin)
			admin.POST("/news", handlers.CreateNews)
			admin.PUT("/news/:id", handlers.UpdateNews)
			admin.DELETE("/news/:id", handlers.DeleteNews)

			// Products
			admin.GET("/products", handlers.GetAllProducts)
			admin.POST("/products", handlers.CreateProduct)
			admin.PUT("/products/:id", handlers.UpdateProduct)
			admin.DELETE("/products/:id", handlers.DeleteProduct)

			// Company Profile
			admin.PUT("/company-profile", handlers.UpdateCompanyProfile)
		}
	}

	return r
}
