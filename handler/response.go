package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"msg":       msg,
		"errorCode": code,
	})
}

func sendSucess(c *gin.Context, op string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  fmt.Sprintf("operation %s successfull", op),
		"data": data,
	})
}
