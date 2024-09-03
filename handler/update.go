package handler

import "github.com/gin-gonic/gin"

func UpdateHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Update",
	})
}
