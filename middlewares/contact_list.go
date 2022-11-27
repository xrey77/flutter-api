package middlewares

import (
	"net/http"
	"src/flutter-api/config"
	"src/flutter-api/models"

	"github.com/gin-gonic/gin"
)

func GetAllContacts(c *gin.Context) {
	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized Access")
		return
	}
	db := config.Connect()
	defer db.Close()

	var contacts models.Contacts
	var jsondata []models.Contacts
	err2 := db.Model(&contacts).Column("*").Select(&jsondata)
	if err2 != nil {
		c.JSON(405, gin.H{"message": err2.Error()})
		return
	} else {
		c.JSON(200, jsondata)
	}
}
