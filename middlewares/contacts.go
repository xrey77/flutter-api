package middlewares

import (
	"encoding/json"
	"fmt"
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

func GetContact(c *gin.Context) {
	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized Access")
		return
	}
	idno := c.Param("id")
	db := config.Connect()
	defer db.Close()

	jsondata := new(models.Contact)
	if err2 := db.Model(jsondata).Where("id = ?", idno).Select(); err2 != nil {
		fmt.Println(err2)
		c.JSON(400, gin.H{"message": err2.Error()})
		return
	} else {
		c.JSON(200, jsondata)
	}
}

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

func DeleteContact(c *gin.Context) {
	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized Access")
		return
	}
	idno := c.Param("id")

	db := config.Connect()
	defer db.Close()

	contactmodel := new(models.Contact)
	_, err = db.Model(contactmodel).Where("id = ?", idno).Delete()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"message": "Contact has been deleted."})
	}
}
