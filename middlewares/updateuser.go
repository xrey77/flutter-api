package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"src/flutter-api/config"
	"src/flutter-api/models"
	"src/flutter-api/utilz"
)

func UpdateUser(c *gin.Context) {
	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized Access")
		return
	}

	//GET PARAM ID
	idno := c.Param("id")
	log.Print("user ID :", idno)
	db := config.Connect()
	defer db.Close()

	//GET INPUT FROM USER
	var user = new(models.TempUsers)
	err2 := json.NewDecoder(c.Request.Body).Decode(user)
	if err2 != nil {
		log.Print("Unable to decode the request body")
		return
	}
	//ENCRYPT PASSWORD
	xbyte := utilz.EncryptPassword(user.Password)
	hashPwd := utilz.HashAndSalt(xbyte)
	user.Password = hashPwd

	var users = new(models.User)

	xid, err3 := strconv.ParseInt(idno, 10, 64)
	if err3 != nil {
		log.Print(err3.Error())
	} else {
		user.ID = xid
		now := time.Now()
		now1 := now.Format("2006-01-02 15:04:05")
		now2, _ := time.Parse("2006-01-02 15:04:05", now1)

		users.Lastname = user.Lastname
		users.Firstname = user.Firstname
		users.Email = user.Email
		users.Mobile = user.Mobile
		users.Username = user.Username
		users.Password = user.Password
		users.Updated_at = now2
		_, err4 := db.Model(users).Where("id = ?", idno).UpdateNotZero()
		if err4 != nil {
			log.Print("Unable to decode the request body.", err.Error())
			c.JSON(http.StatusOK, gin.H{"message": "Unable to decode the request body."})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Record(s) has been updated successfully."})
		}

	}

	// var newFile string
	// if user.Userpicture != "" {
	// 	base64String, _ := base64.StdEncoding.DecodeString(user.Userpicture)

	// 	//CONVERT BASE64 TO STRING
	// 	var data = string(base64String)
	// 	fmt.Println(data)
	// 	idx := strings.Index(data, ";base64,")
	// 	if idx < 0 {
	// 		panic("InvalidImage")
	// 	}

	// 	//IMAGE EXTENSION
	// 	ImageType := data[11:idx]

	// 	unbased, err := base64.StdEncoding.DecodeString(data[idx+8:])
	// 	if err != nil {
	// 		panic("Cannot decode b64")
	// 	}
	// 	//CREATE NEW FILE
	// 	newFile = "./assets/users/00" + idno + "." + ImageType

	// 	r := bytes.NewReader(unbased)
	// 	switch ImageType {
	// 	case "png":

	// 		im, err := png.Decode(r)
	// 		if err != nil {
	// 			panic("Bad png")
	// 		}
	// 		f, err := os.OpenFile(newFile, os.O_WRONLY|os.O_CREATE, 0777)
	// 		if err != nil {
	// 			panic("Cannot open file")
	// 		}
	// 		// Set the expected size, it will reduce to 13kb
	// 		dst := image.NewRGBA(image.Rect(200, 200, im.Bounds().Max.X/2, im.Bounds().Max.Y/2))

	// 		// Resize:
	// 		draw.NearestNeighbor.Scale(dst, dst.Rect, im, im.Bounds(), draw.Over, nil)

	// 		png.Encode(f, dst)

	// 	case "jpeg":

	// 		im, err := jpeg.Decode(r)
	// 		if err != nil {
	// 			panic("Bad jpeg")
	// 		}

	// 		f, err := os.OpenFile(newFile, os.O_WRONLY|os.O_CREATE, 0777)
	// 		if err != nil {
	// 			panic("Cannot open file")
	// 		}

	// 		// Set the expected size, it will reduce to 13kb
	// 		dst := image.NewRGBA(image.Rect(200, 200, im.Bounds().Max.X/2, im.Bounds().Max.Y/2))

	// 		// Resize:
	// 		draw.NearestNeighbor.Scale(dst, dst.Rect, im, im.Bounds(), draw.Over, nil)

	// 		// Encode to `output`:
	// 		jpeg.Encode(f, dst, nil)

	// 	case "gif":

	// 		im, err := gif.Decode(r)
	// 		if err != nil {
	// 			panic("Bad gif")
	// 		}

	// 		f, err := os.OpenFile(newFile, os.O_WRONLY|os.O_CREATE, 0777)
	// 		if err != nil {
	// 			panic("Cannot open file")
	// 		}
	// 		// Set the expected size, it will reduce to 13kb
	// 		dst := image.NewRGBA(image.Rect(200, 200, im.Bounds().Max.X/2, im.Bounds().Max.Y/2))

	// 		// Resize:
	// 		draw.NearestNeighbor.Scale(dst, dst.Rect, im, im.Bounds(), draw.Over, nil)
	// 		gif.Encode(f, dst, nil)

	// 	}
	// } else {
	// 	newFile = "pix.png"
	// }
	// var sqlStatement string
	// var url string
	// var msg string
	// if user.Password != "" {
	// 	xbyte := getPwds(user.Password)
	// 	hashpwd := hashAndSalted(xbyte)
	// 	if user.Userpicture != "" {
	// 		url = "https://golang-api-10.herokuapp.com/" + newFile[2:]
	// 		// url = "http://localhost:8080/" + newFile[2:]
	// 		sqlStatement = `UPDATE USERS SET lastname=$1, firstname=$2, mobile=$3, password=$4, userpicture=$5 WHERE id=$6`
	// 		res, _ := db.Exec(sqlStatement, &user.Lastname, &user.Firstname, &user.Mobile, hashpwd, &url, &idno)
	// 		xidno, _ := res.RowsAffected()
	// 		log.Println(xidno)

	// 	} else {
	// 		if user.Userpicture != "" {
	// 			url = "https://golang-api-10.herokuapp.com/" + newFile[2:]
	// 			// url = "http://localhost:8082/" + newFile[2:]
	// 			sqlStatement = `UPDATE USERS SET lastname=$1, firstname=$2, mobile=$3, password=$4,userpicture=$5 WHERE id=$6`
	// 			res, _ := db.Exec(sqlStatement, &user.Lastname, &user.Firstname, &user.Mobile, hashpwd, &url, &idno)
	// 			xidno, _ := res.RowsAffected()
	// 			log.Println(xidno)

	// 		} else {
	// 			sqlStatement = `UPDATE USERS SET lastname=$1, firstname=$2, mobile=$3, password=$4 WHERE id=$5`
	// 			res, _ := db.Exec(sqlStatement, &user.Lastname, &user.Firstname, &user.Mobile, hashpwd, &idno)
	// 			xidno, _ := res.RowsAffected()
	// 			log.Println(xidno)
	// 		}
	// 	}
	// 	msg = "User ID " + idno + " has been updated."
	// 	c.JSON(http.StatusOK, gin.H{"message": msg})
	// } else {
	// 	if user.Userpicture != "" {
	// 		// url = "http://localhost:8082/" + newFile[2:]
	// 		url = "https://golang-api-10.herokuapp.com/" + newFile[2:]
	// 		sqlStatement = `UPDATE USERS SET lastname=$1, firstname=$2, mobile=$3,userpicture=$4 WHERE id=$5`
	// 		res, _ := db.Exec(sqlStatement, &user.Lastname, &user.Firstname, &user.Mobile, &url, &idno)
	// 		xidno, _ := res.RowsAffected()
	// 		log.Println(xidno)

	// 	} else {
	// 		sqlStatement = `UPDATE USERS SET lastname=$1, firstname=$2, mobile=$3 WHERE id=$4`
	// 		res, _ := db.Exec(sqlStatement, &user.Lastname, &user.Firstname, &user.Mobile, &idno)
	// 		xidno, _ := res.RowsAffected()
	// 		log.Println(xidno)
	// 	}
	// 	msg = "User ID " + idno + " has been updated."
	// 	c.JSON(http.StatusOK, gin.H{"message": msg})
	// }

}

// func UpdateUserMgt(c *gin.Context) {
// 	_, err := ExtractTokenMetadata(c.Request)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, "unauthorized")
// 		return
// 	}
// 	//GET PARAM ID
// 	idno := c.Param("id")
// 	db := config.Connect()
// 	defer db.Close()

// 	//GET INPUT FROM USER
// 	var user models.User
// 	err = json.NewDecoder(c.Request.Body).Decode(&user)
// 	if err != nil {
// 		log.Fatalf("Unable to decode the request body.  %v", err)
// 	}
// 	var sqlStatement string
// 	var msg string
// 	if user.Password != "" {
// 		xbyte := getPwds(user.Password)
// 		hashpwd := hashAndSalted(xbyte)
// 		sqlStatement = `UPDATE USERS SET lastname=$1, firstname=$2, email=$3, mobile=$4, username=$5, password=$6,isactivated=$7 WHERE id=$8`
// 		res, _ := db.Exec(sqlStatement,
// 			&user.Lastname,
// 			&user.Firstname,
// 			&user.Email,
// 			&user.Mobile,
// 			&user.Username,
// 			hashpwd,
// 			&user.Isactivated,
// 			&idno)
// 		xidno, _ := res.RowsAffected()
// 		log.Println(xidno)
// 		msg = "User ID " + idno + " has been updated."
// 		c.JSON(http.StatusOK, gin.H{"message": msg})

// 	} else {

// 		sqlStatement = `UPDATE USERS SET lastname=$1, firstname=$2, email=$3, mobile=$4, username=$5, isactivated=$6 WHERE id=$7`
// 		res, _ := db.Exec(sqlStatement,
// 			&user.Lastname,
// 			&user.Firstname,
// 			&user.Email,
// 			&user.Mobile,
// 			&user.Username,
// 			&user.Isactivated,
// 			&idno)
// 		xidno, _ := res.RowsAffected()
// 		log.Println(xidno)
// 		msg = "User ID " + idno + " has been updated."
// 		c.JSON(http.StatusOK, gin.H{"message": msg})

// 	}
// }

// func getPwds(pwd string) []byte {
// 	return []byte(pwd)
// }

// func hashAndSalted(pwd []byte) string {
// 	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return string(hash)
// }
