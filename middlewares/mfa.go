package middlewares

import (
	"log"
	"net/http"

	"src/flutter-api/globalvar"

	"src/flutter-api/config"
	"src/flutter-api/models"

	"github.com/dgryski/dgoogauth"
	"github.com/gin-gonic/gin"
)

func Mfaverfication(c *gin.Context) {
	mfaToken := c.PostForm("otpMFA")
	uname := globalvar.GetUsername()
	secretBase32 := getsecretkey(uname)

	log.Println("Secret Key 2 : " + secretBase32)
	otpc := &dgoogauth.OTPConfig{
		Secret:      secretBase32,
		WindowSize:  3,
		HotpCounter: 0,
	}

	val, err := otpc.Authenticate(mfaToken)
	if err != nil {
		c.JSON(200, gin.H{"statuscode": 400, "message": err.Error()})
		return
	}

	if !val {
		c.JSON(200, gin.H{"statuscode": 400, "message": "OTP Token is not valid."})
		return
	}
	c.JSON(200, gin.H{"statuscode": 200, "message": "MFA Verfication successfull."})
	c.Redirect(http.StatusFound, "/")
}

func getsecretkey(uname string) string {
	db := config.Connect()
	defer db.Close()
	var tmpUser models.TempUsers
	// FIND USERNAME IF IT IS ALREADY EXISTS
	sqlFindUsername := `SELECT secretkey FROM USERS WHERE username=$1;`
	rowUsername := db.QueryRow(sqlFindUsername, uname)
	err2 := rowUsername.Scan(&tmpUser.Secretkey)
	if err2 != nil {
		return ""
	} else {
		return tmpUser.Secretkey
	}
}
