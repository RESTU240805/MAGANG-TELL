package routes

import (
	"time"

	"magang-unpra-backend/handlers"
	"magang-unpra-backend/middleware"

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
		api.GET("/about-section", handlers.GetAboutSection)
		api.PUT("/about-section", handlers.UpdateAboutSection)
		api.GET("/product-slides", handlers.GetAllSlides)
		api.GET("/product-page", handlers.GetProductPage)
		api.PUT("/product-page", handlers.UpdateProductPage)
		api.GET("/team-members", handlers.GetAllTeamMembers)

		// Admin
		admin := api.Group("/admin")
		admin.Use(middleware.AuthRequired())
		{
			admin.GET("/news", handlers.GetAllNewsAdmin)
			admin.POST("/news", handlers.CreateNews)
			admin.PUT("/news/:id", handlers.UpdateNews)
			admin.DELETE("/news/:id", handlers.DeleteNews)

			admin.GET("/products", handlers.GetAllProducts)
			admin.POST("/products", handlers.CreateProduct)
			admin.PUT("/products/:id", handlers.UpdateProduct)
			admin.DELETE("/products/:id", handlers.DeleteProduct)

			admin.GET("/product-slides", handlers.GetAllSlides)
			admin.POST("/product-slides", handlers.CreateSlide)
			admin.PUT("/product-slides/:id", handlers.UpdateSlide)
			admin.DELETE("/product-slides/:id", handlers.DeleteSlide)

			admin.PUT("/company-profile", handlers.UpdateCompanyProfile)
			admin.GET("/about-section", handlers.GetAboutSection)
			admin.PUT("/about-section", handlers.UpdateAboutSection)
			admin.GET("/product-page", handlers.GetProductPage)
			admin.PUT("/product-page", handlers.UpdateProductPage)

			admin.GET("/team-members", handlers.GetAllTeamMembersAdmin)
			admin.POST("/team-members", handlers.CreateTeamMember)
			admin.PUT("/team-members/:id", handlers.UpdateTeamMember)
			admin.DELETE("/team-members/:id", handlers.DeleteTeamMember)
		}
	}

	return r
}
