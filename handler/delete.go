package handler

import (
	"fmt"
	"net/http"

	"github.com/CarlosHoff/crud.git/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteHandler(c *gin.Context) {

	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	crud := schemas.Crud{}

	if err := db.First(&crud, id).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("Crud with ID: %s not found", id))
		return
	}

	if err := db.Delete(&crud, id).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("Error to delete ID: %s, try again later", id))
		return
	}

	sendSucess(c, "delete-crud", crud)

}
