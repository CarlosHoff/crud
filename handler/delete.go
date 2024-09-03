package handler

import "github.com/gin-gonic/gin"

func DeleteHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete",
	})
}
