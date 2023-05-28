package controllers

import (
	"github.com/diiineeei/IntegracaoJson/authjwt"
	"github.com/diiineeei/IntegracaoJson/database"
	"github.com/diiineeei/IntegracaoJson/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterContactsList(context *gin.Context) {
	company, err := authjwt.ValidateToken(context.GetHeader("Authorization"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	var contacts models.Contacts
	if err := context.ShouldBindJSON(&contacts); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	switch company {
	case models.VAREJAO:
		for _, con := range contacts.Contacts {
			con.Company = company
			con.NormalizeContacts()
			record := database.InstancePostgres.Create(&con)
			if record.Error != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
				context.Abort()
				return
			}
		}
	case models.MACAPA:
		for _, con := range contacts.Contacts {
			con.Company = company
			con.NormalizeContacts()
			record := database.InstanceMySQL.Create(&con)
			if record.Error != nil {
				context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
				context.Abort()
				return
			}
		}
	default:
		context.JSON(http.StatusBadRequest, gin.H{"error": "Company not found"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"info": "Created"})
}

func RegisterContacts(context *gin.Context) {
	company, err := authjwt.ValidateToken(context.GetHeader("Authorization"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	var contact models.Contact
	if err := context.ShouldBindJSON(&contact); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	contact.Company = company
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
