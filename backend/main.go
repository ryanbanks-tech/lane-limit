package main

import (
	"log"

	"lane-limit/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Configure CORS if needed (for local dev, etc.)
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))

	// Set up routes
	api := r.Group("/api")
	{
		// Lanes endpoints
		api.GET("/lanes", controllers.GetAllLanes)
		api.POST("/lanes", controllers.UpsertLane)
		api.DELETE("/lanes/:id", controllers.ClearLaneInfo)
		api.DELETE("/lanes/clearAll", controllers.ClearAllLanes)

		// Image upload
		api.POST("/image", controllers.UploadImage)
	}

	log.Println("Starting server on :8080...")
	// Run the server
	r.Run(":8080")
}
