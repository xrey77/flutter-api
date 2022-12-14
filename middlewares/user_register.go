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

	var usermodel = new(models.User)
	err := json.NewDecoder(c.Request.Body).Decode(&usermodel)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	c.Header("Content-Type", "application/json;charset=UTF-8")
	email := usermodel.Email
	username := usermodel.Username
	pwd := usermodel.Password

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
	usermodel.ID = cnt
	usermodel.Password = hashPwd
	usermodel.Userpicture = urlpic
	usermodel.Secretkey = ""
	usermodel.Qrcode = ""

	//INSERT DATA
	_, err3 := db.Model(usermodel).Insert()
	if err3 != nil {
		// c.JSON(500, gin.H{"message 2": err3.Error()})
		c.JSON(500, gin.H{"message": "Unable to execute the query."})
		return
	} else {
		msg := "You have registered successfully."
		c.JSON(201, gin.H{"message": msg})
	}
}
