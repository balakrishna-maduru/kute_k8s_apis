package routes

import (
	"kute_k8s_apis/internal/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the routes for the application.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/api/v1/namespaces", controllers.GetNamespaces)
	r.POST("/api/v1/namespaces", controllers.CreateNamespace)
	r.GET("/api/v1/namespaces/:name", controllers.GetNamespace)
	r.DELETE("/api/v1/namespaces/:name", controllers.DeleteNamespace)
	return r
}

// RunServer starts the HTTP server.
func RunServer() {
	r := SetupRouter()
	r.Run("localhost:8080")
}
