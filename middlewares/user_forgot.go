package middlewares

import (
	"log"
	"net/http"

	"src/flutter-api/config"
	"src/flutter-api/models"
	"src/flutter-api/utilz"

	"github.com/gin-gonic/gin"
)

func Forgot(c *gin.Context) {
	c.Header("Content-Type", "application/json charset=utf-8")
	xmail := c.PostForm("xmail")
	if utilz.ValidateEmail(xmail) == 0 {
		c.JSON(200, gin.H{"statuscode": 404, "message": "Email Address not found."})
		return
	}
	mtoken := utilz.RandomToken()
	msg := "<div>Please copy and paste Mail Token below:</div></br><br/><div style=\"font-size:24px;\">" + mtoken + "</div>"
	SendMail(msg, "Mail Token", xmail)

	db := config.Connect()
	defer db.Close()
	var users = new(models.User)
	users.Mailtoken = mtoken
	_, err4 := db.Model(users).Where("email = ?", xmail).UpdateNotZero()
	if err4 != nil {
		log.Print("Unable to decode the request body.", err4.Error())
		c.JSON(http.StatusOK, gin.H{"message": "Unable to decode the request body."})
		return
	} else {
		c.JSON(200, gin.H{"statuscode": 200, "message": "Mail Token has been sent to " + xmail})
	}
}
