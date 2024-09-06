package handler

import (
	"net/http"

	"github.com/CarlosHoff/crud.git/schemas"
	"github.com/gin-gonic/gin"
)

func SearchHandler(c *gin.Context) {

	crud := []schemas.Crud{}

	if err := db.Find(&crud).Error; err != nil {
		sendError(c, http.StatusNoContent, "No Registers ind DB")
		return
	}

	sendSucess(c, "search-crud", crud)

}
