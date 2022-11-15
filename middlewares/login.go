package middlewares

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"src/flutter-api/config"
	"src/flutter-api/models"
	"src/flutter-api/utilz"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Login(c *gin.Context) {

	var user models.UserLogin
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	c.Header("Content-Type", "application/json charset=utf-8")
	uname := user.UserName
	pwd2 := user.PassWord

	if utilz.ValidateUserName(uname) == 0 {
		c.JSON(400, gin.H{"message": "Username is not yet registered, please click register link at the top right of the screen."})
		return
	}

	if utilz.IsActivated(uname) == 0 {
		c.JSON(200, gin.H{"message": "You account is not yet activated, please check you Email inbox and click Activate button."})
		return
	}

	db := config.Connect()
	defer db.Close()
	var tmpUser models.UserLogin
	sqlFindUsername := `SELECT id,username,password,userpicture,role,otp FROM USERS WHERE username=$1;`
	rowUsername := db.QueryRow(sqlFindUsername, uname)
	err2 := rowUsername.Scan(&tmpUser.ID, &tmpUser.UserName, &tmpUser.PassWord, &tmpUser.Userpicture, &tmpUser.Role, &tmpUser.Otp)
	if err2 != nil {
		fmt.Println(err2)
	}

	hPwd := tmpUser.PassWord

	if utilz.ComparePassword(hPwd, utilz.EncryptPassword(pwd2)) {
		xid := utilz.GetUID(uname)
		ids, err := strconv.ParseInt(xid, 10, 64)
		if err != nil {
			panic(err)
		}

		tk := &models.Token{UserId: uint(ids)}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		tokenString, _ := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
		tmpUser.Token = tokenString
		totp := strconv.FormatInt(tmpUser.Otp, 2)
		msg := map[string]string{"id": xid, "username": tmpUser.UserName, "token": tmpUser.Token, "userpicture": tmpUser.Userpicture, "role": tmpUser.Role, "otp": totp}
		c.JSON(http.StatusOK, msg)
	} else {
		c.JSON(200, gin.H{"message": "Wrong password."})
		return
	}
}
