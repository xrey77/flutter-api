package middlewares

import (
	"fmt"
	"net/http"

	"src/flutter-api/models"

	"src/flutter-api/config"

	"github.com/gin-gonic/gin"
)

func GetUserbyID(c *gin.Context) {

	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	id := c.Param("id")
	db := config.Connect()
	defer db.Close()

	var tmpUser models.TempUsers
	sqlFindUser := `SELECT id,lastname, firstname, email, mobile, username, userpicture, secretkey, qrcode, otp, role FROM USERS WHERE id=$1;`

	rowUsername := db.QueryRow(sqlFindUser, &id)
	err2 := rowUsername.Scan(
		&tmpUser.ID,
		&tmpUser.Lastname,
		&tmpUser.Firstname,
		&tmpUser.Email,
		&tmpUser.Mobile,
		&tmpUser.Username,
		&tmpUser.Userpicture,
		&tmpUser.Secretkey,
		&tmpUser.Qrcode,
		&tmpUser.Otp,
		&tmpUser.Role)
	if err2 != nil {
		fmt.Println(err2)
		c.JSON(400, gin.H{"message": err2.Error()})
		return
	} else {
		c.JSON(200, tmpUser)

	}

}

func GetUser(c *gin.Context) {

	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	id := c.Param("id")
	db := config.Connect()
	defer db.Close()

	var tmpUser models.TempUsers
	sqlFindUser := `SELECT id,lastname, firstname, email, mobile, username, password, isactivated, role FROM USERS WHERE id=$1;`

	rowUsername := db.QueryRow(sqlFindUser, &id)
	err2 := rowUsername.Scan(
		&tmpUser.ID,
		&tmpUser.Lastname,
		&tmpUser.Firstname,
		&tmpUser.Email,
		&tmpUser.Mobile,
		&tmpUser.Username,
		&tmpUser.Password,
		&tmpUser.Isactivated,
		&tmpUser.Role)
	if err2 != nil {
		// fmt.Println(err2)
		c.JSON(400, gin.H{"message": err2.Error()})
		return
	} else {
		c.JSON(200, tmpUser)

	}

}
