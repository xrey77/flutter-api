package middlewares

import (
	"encoding/json"
	"net/http"
	"src/flutter-api/config"
	"src/flutter-api/models"

	"github.com/gin-gonic/gin"
)

func UpdateContact(c *gin.Context) {
	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized Access")
		return
	}

	idno := c.Param("id")
	var contact models.Contact
	err = json.NewDecoder(c.Request.Body).Decode(&contact)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	db := config.Connect()
	defer db.Close()

	contacts := new(models.Contact)
	contacts.Contact_name = contact.Contact_name
	contacts.Contact_address = contact.Contact_address
	contacts.Contact_email = contact.Contact_email
	contacts.Contact_mobileno = contact.Contact_mobileno
	contacts.Isactive = contact.Isactive
	_, err4 := db.Model(contacts).Where("id = ?", idno).UpdateNotZero()
	if err4 != nil {
		c.JSON(400, gin.H{"message": "Unable to decode the request body."})
		return
	} else {
		c.JSON(200, gin.H{"message": "Record(s) has been updated successfully."})
	}
}
