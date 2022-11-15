package middlewares

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"src/flutter-api/models"

	"src/flutter-api/config"

	"github.com/gin-gonic/gin"
)

func AddContact(c *gin.Context) {
	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	var contact models.Contact
	err = json.NewDecoder(c.Request.Body).Decode(&contact)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	if validateContact(contact.Contact_Name) != 0 {
		c.JSON(400, gin.H{"message": "Contact Name is taken."})
		return
	}

	fmt.Println(contact)
	db := config.Connect()
	defer db.Close()

	sqlStatement := `INSERT INTO CONTACTS(id, contact_name, contact_address, contact_email, contact_mobileno, isactive) VALUES(nextval('contactid'),$1, $2, $3, $4, $5)`
	res, err3 := db.Exec(
		sqlStatement,
		&contact.Contact_Name,
		&contact.Contact_Address,
		&contact.Contact_Email,
		&contact.Contact_Mobileno,
		&contact.IsActive)
	if err3 != nil {
		c.JSON(400, gin.H{"message": "Unable to execute the query."})
		return
	}
	idno, _ := res.RowsAffected()
	fmt.Println(idno)
	msg := "Contact Name, " + contact.Contact_Name + " has been added."
	c.JSON(200, gin.H{"message": msg})
}

func GetAllContacts(c *gin.Context) {
	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	var contact []models.Contact
	db := config.Connect()
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 55*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT id,contact_name,contact_address,contact_email,contact_mobileno,isactive FROM contacts ORDER BY id")
	if err != nil {
		c.JSON(400, gin.H{"message": err})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var contacts models.Contact
		rows.Scan(&contacts.ID,
			&contacts.Contact_Name,
			&contacts.Contact_Address,
			&contacts.Contact_Email,
			&contacts.Contact_Mobileno,
			&contacts.IsActive)
		contact = append(contact, contacts)
	}
	c.JSON(200, contact)

}

func GetContact(c *gin.Context) {
	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	idno := c.Param("id")
	var contact models.Contact
	db := config.Connect()
	defer db.Close()
	sqlFindContact := `SELECT id,contact_name,contact_address,contact_email, contact_mobileno,isactive FROM CONTACTS WHERE id=$1;`

	rowUsername := db.QueryRow(sqlFindContact, &idno)
	err2 := rowUsername.Scan(
		&contact.ID,
		&contact.Contact_Name,
		&contact.Contact_Address,
		&contact.Contact_Email,
		&contact.Contact_Mobileno,
		&contact.IsActive)
	if err2 != nil {
		fmt.Println(err2)
		c.JSON(400, gin.H{"message": err2.Error()})
		return
	} else {
		c.JSON(200, contact)

	}

}

func UpdateContact(c *gin.Context) {
	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	idno := c.Param("id")
	var contact models.Contact
	err = json.NewDecoder(c.Request.Body).Decode(&contact)
	if err != nil {
		c.JSON(400, gin.H{"message": err})
		return
	}
	db := config.Connect()
	defer db.Close()

	sqlStatement := `UPDATE CONTACTS SET contact_name=$1, contact_address=$2, contact_email=$3, contact_mobileno=$4,isactive=$5 WHERE id=$6`
	res, err1 := db.Exec(sqlStatement,
		&contact.Contact_Name,
		&contact.Contact_Address,
		&contact.Contact_Email,
		&contact.Contact_Mobileno,
		&contact.IsActive,
		&idno)
	if err1 != nil {
		c.JSON(400, gin.H{"message": err1})
		return
	}
	xidno, _ := res.RowsAffected()
	if xidno == 0 {
		c.JSON(400, gin.H{"message": "Contact ID does not exits."})
		return
	} else {
		msg := "Contact ID, " + idno + " has been updated."
		c.JSON(200, gin.H{"message": msg})
	}

}

func DeleteContact(c *gin.Context) {
	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	idno := c.Param("id")

	db := config.Connect()
	defer db.Close()

	query := "DELETE FROM CONTACTS WHERE id = $1"

	var rowsAffected int64
	var result sql.Result

	result, err = db.Exec(query, idno)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	if rowsAffected > 0 {
		msg := map[string]string{"msg": "Contact ID, " + idno + " has been deleted."}
		c.JSON(http.StatusFound, gin.H{"message": msg})
		return
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Contact ID, " + idno + " does not exists."})
	}

}

func validateContact(cname string) int64 {
	db := config.Connect()
	defer db.Close()
	var tmpContact models.Contact
	sqlFindContactname := `SELECT contact_name FROM contacts WHERE contact_name=$1;`
	rowContactname := db.QueryRow(sqlFindContactname, cname)
	err2 := rowContactname.Scan(&tmpContact.Contact_Name)
	if err2 != nil {
		return 0
	} else {
		return 1
	}
}
