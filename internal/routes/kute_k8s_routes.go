// internal/routes/routes.go

package routes

import (
	"kute_k8s_apis/internal/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the routes for the application.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Define routes
	r.GET("/albums", controllers.GetAlbums)
	r.GET("/albums/:id", controllers.GetAlbumByID)
	r.POST("/albums", controllers.PostAlbums)

	return r
}

// RunServer starts the HTTP server.
func RunServer() {
	r := SetupRouter()
	r.Run("localhost:8080")
}
