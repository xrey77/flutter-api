package middlewares

import (
	"log"
	"math/rand"
	"time"

	"src/flutter-api/config"
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
	mtoken := randomToken()
	msg := "<div>Please copy and paste Mail Token below:</div></br><br/><div style=\"font-size:24px;\">" + mtoken + "</div>"
	SendMail(msg, "Mail Token", xmail)

	db := config.Connect()
	defer db.Close()
	sqlStatement := `UPDATE USERS SET mailtoken=$1 WHERE email=$2`
	res, _ := db.Exec(sqlStatement, &mtoken, &xmail)
	idno, _ := res.RowsAffected()
	log.Println(idno)
	c.JSON(200, gin.H{"statuscode": 200, "message": "Mail Token has been sent to " + xmail})
}

func randomToken() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var codes [6]byte
	for i := 0; i < 6; i++ {
		codes[i] = uint8(48 + r.Intn(10))
	}
	return string(codes[:])
}
