package controllers

import (
	"github.com/diiineeei/IntegracaoJson/authjwt"
	"github.com/diiineeei/IntegracaoJson/database"
	"github.com/diiineeei/IntegracaoJson/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterContacts(context *gin.Context) {
	company, err := authjwt.ValidateToken(context.GetHeader("Authorization"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	var contact models.Contacts
	contact.Company = company
	if err := context.ShouldBindJSON(&contact); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	
	contact.NormalizeContacts()

	switch company {
	case models.VAREJAO:
		record := database.InstancePostgres.Create(&contact)
		if record.Error != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
			context.Abort()
			return
		}
	case models.MACAPA:
		record := database.InstanceMySQL.Create(&contact)
		if record.Error != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
			context.Abort()
			return
		}
	}

	context.JSON(http.StatusCreated, gin.H{"userId": contact.ID, "Name": contact.Name, "CellPhone": contact.CellPhone, "Company": contact.Company})
}
