package router

import (
	"github.com/CarlosHoff/crud.git/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {

	v1 := r.Group("/api/v1")

	{
		v1.GET("/health", func(c *gin.Context) {
			handler.HealthHandler(c)
		})
	}

	{
		v1.POST("/create", func(c *gin.Context) {
			handler.CreateHandler(c)
		})
	}

	{
		v1.DELETE("/delete", func(c *gin.Context) {
			handler.DeleteHandler(c)
		})
	}

	{
		v1.PUT("/update", func(c *gin.Context) {
			handler.UpdateHandler(c)
		})
	}

	{
		v1.GET("/search", func(c *gin.Context) {
			handler.SearchHandler(c)
		})
	}

}
