package router

import "github.com/gin-gonic/gin"

func initializeRoutes(r *gin.Engine) {

	v1 := r.Group("/api/v1")

	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Projeto UP",
			})
		})
	}

	{
		v1.POST("/create", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Projeto UP",
			})
		})
	}

	{
		v1.DELETE("/delete", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Projeto UP",
			})
		})
	}

	{
		v1.PUT("/update", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Projeto UP",
			})
		})
	}

	{
		v1.GET("/search", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Projeto UP",
			})
		})
	}

}
