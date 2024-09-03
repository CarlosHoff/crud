package handler

import "github.com/gin-gonic/gin"

func SearchHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Search",
	})
}
