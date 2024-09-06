package handler

import (
	"fmt"
	"net/http"

	"github.com/CarlosHoff/crud.git/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateHandler(c *gin.Context) {

	request := struct {
		CreateRequest
	}{}

	c.BindJSON(&request)

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

	crud.Role = request.Role
	crud.Company = request.Company
	crud.Location = request.Location
	crud.Remote = *request.Remote
	crud.Link = request.Link
	crud.Salary = request.Salary

	if err := db.Model(&crud).Updates(map[string]interface{}{
		"role":     crud.Role,
		"company":  crud.Company,
		"location": crud.Location,
		"remote":   crud.Remote,
		"link":     crud.Link,
		"salary":   crud.Salary,
	}).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("Error updating Crud with ID: %s", id))
		return
	}

	sendSucess(c, "update-crud", crud)

}
