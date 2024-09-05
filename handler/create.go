package handler

import "github.com/gin-gonic/gin"

func CreateHandler(c *gin.Context) {
	request := struct {
		CreateRequest
	}{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("Erro ao validar os campos: %v", err.Error())
		return
	}

	if err := db.Create(&request).Error; err != nil {
		logger.Errf("Erro ao criar Crud: %v", err.Error())
		return
	}
}
