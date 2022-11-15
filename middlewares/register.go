package middlewares

import (
	"encoding/json"
	"fmt"
	"log"

	"src/flutter-api/models"
	"src/flutter-api/utilz"

	"src/flutter-api/config"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	c.Header("Content-Type", "application/json charset=utf-8")
	lastname := user.Lastname
	firstname := user.Firstname
	email := user.Email
	mobile := user.Mobile
	username := user.Username
	pwd := user.Password

	if utilz.ValidateEmail(email) > 0 {
		c.JSON(200, gin.H{"message": "Email has already taken."})
		return
	}

	if utilz.ValidateUserName(username) > 0 {
		c.JSON(201, gin.H{"message": "Username has already taken."})
		return
	}

	db := config.Connect()
	defer db.Close()
	//ENCRYPT PASSWORD
	xbyte := utilz.EncryptPassword(pwd)
	hashPwd := utilz.HashAndSalt(xbyte)
	//INSERT DATA
	urlpic := "http://localhost:9000/assets/users/pix.png"

	// urlpic := "https://golang-api-10.herokuapp.com/assets/users/pix.png"
	sqlStatement := `INSERT INTO users(
		id, 
		lastname, 
		firstname, 
		email, 
		username, 
		password, 
		isactivated,
		mobile, 
		userpicture, 
		secretkey,
		qrcode) VALUES(nextval('userid'),$1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	res, err3 := db.Exec(sqlStatement, &lastname, &firstname, &email, &username, &hashPwd, 1, &mobile, &urlpic, "", "")
	if err3 != nil {
		c.JSON(400, gin.H{"message": "Unable to execute the query."})
		return
	}
	idno, _ := res.RowsAffected()
	fmt.Println(idno)
	msg := "You have registered successfully."
	c.JSON(200, gin.H{"message": msg})

}
