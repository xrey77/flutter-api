package middlewares

import (
	"encoding/json"
	"log"
	"net/http"

	"src/flutter-api/models"
	"src/flutter-api/utilz"

	"src/flutter-api/config"

	"github.com/gin-gonic/gin"
)

func AddContact(c *gin.Context) {
	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized Access")
		return
	}

	var contact models.Contact
	err = json.NewDecoder(c.Request.Body).Decode(&contact)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	if utilz.ValidateContact(contact.Contact_name) != 0 {
		c.JSON(400, gin.H{"message": "Contact Name is already taken."})
		return
	} else {
		//INSERT DATA
		db := config.Connect()
		defer db.Close()
		contacts := new(models.Contact)
		contacts.ID = utilz.GetCID()
		contacts.Contact_name = contact.Contact_name
		contacts.Contact_address = contact.Contact_address
		contacts.Contact_email = contact.Contact_email
		contacts.Contact_mobileno = contact.Contact_mobileno
		contacts.Isactive = contact.Isactive
		_, err3 := db.Model(contacts).Insert()
		if err3 != nil {
			c.JSON(500, gin.H{"message": "Unable to execute the query."})
			return
		} else {
			msg := "Contact has been added successfully."
			c.JSON(201, gin.H{"message": msg})
		}

	}
}
