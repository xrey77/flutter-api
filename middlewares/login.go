package middlewares

import (
	"encoding/json"
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
	var usermodel = new(models.TmpLogin)
	err := json.NewDecoder(c.Request.Body).Decode(usermodel)
	if err != nil {
		log.Print("Unable to decode the request body")
		return
	}

	c.Header("Content-Type", "application/json;charset=UTF-8")

	uname := usermodel.UserName
	pwd2 := usermodel.PassWord
	vUsername := utilz.ValidateUserName(uname)
	if (vUsername) == 0 {
		c.JSON(404, gin.H{"message": "Username is not yet registered, please click register link at the top right of the screen."})
		return
	} else {
		if utilz.IsActivated(uname) == 0 {
			c.JSON(406, gin.H{"message": "You account is not yet activated, please check you Email inbox and click Activate button."})
			return
		} else {
		}
	}

	db := config.Connect()
	defer db.Close()
	var usersmodel = new(models.Users)
	if err2 := db.Model(usersmodel).Where("username = ?", uname).Select(); err2 != nil {
		log.Print(err2.Error())
		return
	} else {
		hPwd := usersmodel.Password
		if utilz.ComparePassword(hPwd, utilz.EncryptPassword(pwd2)) {
			xid := usersmodel.ID
			tk := &models.Token{UserId: uint(xid)}
			token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
			tokenString, _ := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
			totp := strconv.FormatInt(usersmodel.Otp, 2)
			isActive := strconv.FormatInt(usersmodel.Isactivated, 2)
			msg := map[string]string{"id": strconv.Itoa(int(xid)), "username": usersmodel.Username, "token": tokenString, "userpicture": usersmodel.Userpicture, "role": usersmodel.Role, "otp": totp, "isactivated": isActive, "email": usersmodel.Email}
			c.JSON(http.StatusOK, msg)

		} else {
			c.JSON(401, gin.H{"message": "Wrong password."})
			return
		}

	}
}
