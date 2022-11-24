package middlewares

import (
	"encoding/json"
	"log"
	"src/flutter-api/config"
	"src/flutter-api/models"
	"src/flutter-api/utilz"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	cnt := utilz.GetNextval()

	var user = new(models.User)
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	c.Header("Content-Type", "application/json;charset=UTF-8")
	email := user.Email
	username := user.Username
	pwd := user.Password

	if utilz.ValidateEmail(email) > 0 {
		c.JSON(405, gin.H{"message": "Email has already taken."})
		return
	} else {
		if utilz.ValidateUserName(username) > 0 {
			c.JSON(406, gin.H{"message": "Username has already taken."})
			return
		} else {
		}
	}
	db := config.Connect()
	defer db.Close()

	//ENCRYPT PASSWORD
	xbyte := utilz.EncryptPassword(pwd)
	hashPwd := utilz.HashAndSalt(xbyte)

	urlpic := "http://localhost:9000/assets/users/pix.png"
	user.ID = cnt
	user.Password = hashPwd
	user.Userpicture = urlpic
	user.Secretkey = ""
	user.Qrcode = ""

	//INSERT DATA
	_, err3 := db.Model(user).Insert()
	if err3 != nil {
		// c.JSON(500, gin.H{"message 2": err3.Error()})
		c.JSON(500, gin.H{"message": "Unable to execute the query."})
		return
	} else {
		msg := "You have registered successfully."
		c.JSON(200, gin.H{"message": msg})
	}
}
