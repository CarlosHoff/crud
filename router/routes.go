package router

import (
	"github.com/CarlosHoff/crud.git/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	handler.InitializeHandler()
	v1 := r.Group("/api/v1")

	{
		v1.GET("/health", handler.HealthHandler)
		v1.POST("/create", handler.CreateHandler)
		v1.DELETE("/delete", handler.DeleteHandler)
		v1.PUT("/update", handler.UpdateHandler)
		v1.GET("/search", handler.SearchHandler)
	}

}
