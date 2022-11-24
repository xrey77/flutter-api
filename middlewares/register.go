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
	// lastname := user.Lastname
	// firstname := user.Firstname
	email := user.Email
	// mobile := user.Mobile
	username := user.Username
	pwd := user.Password

	// log.Print(lastname)
	// log.Print(firstname)
	// log.Print(email)
	// log.Print(mobile)
	// log.Print(username)
	// log.Print(pwd)
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
	// log.Print(hashPwd)
	//INSERT DATA
	urlpic := "http://localhost:9000/assets/users/pix.png"
	// int64 xidno = nextval('userid')
	user.ID = cnt
	user.Password = hashPwd
	user.Userpicture = urlpic
	user.Secretkey = ""
	user.Qrcode = ""
	// urlpic := "https://golang-api-10.herokuapp.com/assets/users/pix.png"
	// sql := `INSERT INTO USERS(
	// 	id,
	// 	lastname,
	// 	firstname,
	// 	email,
	// 	username,
	// 	password,
	// 	isactivated,
	// 	mobile,
	// 	userpicture,
	// 	secretkey,
	// 	qrcode) VALUES(nextval('userid'),$1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	// _, err3 := db.Exec(sql, &lastname, &firstname, &email, &username, &hashPwd, 1, &mobile, &urlpic, "", "")
	_, err3 := db.Model(user).Insert()
	if err3 != nil {
		c.JSON(500, gin.H{"message 2": err3.Error()})
		// c.JSON(500, gin.H{"message": "Unable to execute the query."})
		return
	} else {
		// idno, _ := res.RowsAffected()
		// fmt.Println(idno)
		msg := "You have registered successfully."
		c.JSON(200, gin.H{"message": msg})
	}
}
