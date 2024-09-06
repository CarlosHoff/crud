package handler

import (
	"net/http"

	"github.com/CarlosHoff/crud.git/schemas"
	"github.com/gin-gonic/gin"
)

func CreateHandler(c *gin.Context) {
	request := struct {
		CreateRequest
	}{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("Erro ao validar os campos: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	crudRecord := schemas.Crud{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&crudRecord).Error; err != nil {
		logger.Errf("Erro ao criar Crud: %v", err.Error())
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucess(c, "create-crud", crudRecord)
}
