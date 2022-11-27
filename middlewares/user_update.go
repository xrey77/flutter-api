package middlewares

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/image/draw"

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
		c.JSON(404, "Unable to decode the request body.")
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

		if user.Password != "" {
			users.Password = hashPwd
		}
		users.Lastname = user.Lastname
		users.Firstname = user.Firstname
		users.Email = user.Email
		users.Mobile = user.Mobile
		users.Updated_at = now2

		//============ START USER PICTURE ==================

		var newFile string
		if user.Userpicture != "" {
			base64String, _ := base64.StdEncoding.DecodeString(user.Userpicture)

			//CONVERT BASE64 TO STRING
			var data = string(base64String)
			fmt.Println(data)
			idx := strings.Index(data, ";base64,")
			if idx < 0 {
				panic("InvalidImage")
			}

			//IMAGE EXTENSION
			ImageType := data[11:idx]

			unbased, err := base64.StdEncoding.DecodeString(data[idx+8:])
			if err != nil {
				panic("Cannot decode b64")
			}
			// CREATE NEW FILE
			newFile = "./assets/users/00" + idno + "." + ImageType

			r := bytes.NewReader(unbased)
			switch ImageType {
			case "png":

				im, err := png.Decode(r)
				if err != nil {
					panic("Bad png")
				}
				f, err := os.OpenFile(newFile, os.O_WRONLY|os.O_CREATE, 0777)
				if err != nil {
					panic("Cannot open file")
				}
				// SET THE EXPECTED SIZE, REDUCE TO 13KB
				dst := image.NewRGBA(image.Rect(200, 200, im.Bounds().Max.X/2, im.Bounds().Max.Y/2))

				// RESIZE
				draw.NearestNeighbor.Scale(dst, dst.Rect, im, im.Bounds(), draw.Over, nil)

				// ENCODE TO OUTPUT
				png.Encode(f, dst)

			case "jpeg":

				im, err := jpeg.Decode(r)
				if err != nil {
					panic("Bad jpeg")
				}

				f, err := os.OpenFile(newFile, os.O_WRONLY|os.O_CREATE, 0777)
				if err != nil {
					panic("Cannot open file")
				}

				// SET THE EXPECTED SIZE, REDUCE TO 13KB
				dst := image.NewRGBA(image.Rect(200, 200, im.Bounds().Max.X/2, im.Bounds().Max.Y/2))

				// RESIZE
				draw.NearestNeighbor.Scale(dst, dst.Rect, im, im.Bounds(), draw.Over, nil)

				// ENCODE TO OUTPUT
				jpeg.Encode(f, dst, nil)

			case "gif":

				im, err := gif.Decode(r)
				if err != nil {
					panic("Bad gif")
				}

				f, err := os.OpenFile(newFile, os.O_WRONLY|os.O_CREATE, 0777)
				if err != nil {
					panic("Cannot open file")
				}
				// SET THE EXPECTED SIZE, REDUCE TO 13KB
				dst := image.NewRGBA(image.Rect(200, 200, im.Bounds().Max.X/2, im.Bounds().Max.Y/2))

				// RESIZE
				draw.NearestNeighbor.Scale(dst, dst.Rect, im, im.Bounds(), draw.Over, nil)

				//ENCODE TO OUTPUT
				gif.Encode(f, dst, nil)

			}
		} else {
			newFile = "pix.png"
		}
		if user.Userpicture != "" {
			// url = "https://golang-api-10.herokuapp.com/" + newFile[2:]
			var url = "http://localhost:9000/" + newFile[2:]
			users.Userpicture = url
		}

		//========= END USER PICTURE =======================

		_, err4 := db.Model(users).Where("id = ?", idno).UpdateNotZero()
		if err4 != nil {
			log.Print("Unable to decode the request body.", err.Error())
			c.JSON(http.StatusOK, gin.H{"message": "Unable to decode the request body."})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Record(s) has been updated successfully."})
		}

	}
}

func UserManagement(c *gin.Context) {
	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized Access")
		return
	}

	//GET USER ID PARAM
	idno := c.Param("id")
	db := config.Connect()
	defer db.Close()

	//GET INPUT FROM USER
	var user = new(models.TempUsers)
	err2 := json.NewDecoder(c.Request.Body).Decode(&user)
	if err2 != nil {
		c.JSON(404, "Unable to decode the request body.")
		return
	}

	// CURRENT UPDATE DATE TIME
	now := time.Now()
	now1 := now.Format("2006-01-02 15:04:05")
	now2, _ := time.Parse("2006-01-02 15:04:05", now1)

	var users = new(models.User)

	xbyte := utilz.EncryptPassword(user.Password)
	hashpwd := utilz.HashAndSalt(xbyte)
	if user.Password != "" {
		users.Password = hashpwd
	}

	// IF THERE IS USER PICTURE
	var newFile string
	if user.Userpicture != "" {
		base64String, _ := base64.StdEncoding.DecodeString(user.Userpicture)

		//CONVERT BASE64 TO STRING
		var data = string(base64String)
		fmt.Println(data)
		idx := strings.Index(data, ";base64,")
		if idx < 0 {
			panic("InvalidImage")
		}

		//IMAGE EXTENSION
		ImageType := data[11:idx]

		unbased, err := base64.StdEncoding.DecodeString(data[idx+8:])
		if err != nil {
			panic("Cannot decode b64")
		}
		// CREATE NEW FILE
		newFile = "./assets/users/00" + idno + "." + ImageType

		r := bytes.NewReader(unbased)
		switch ImageType {
		case "png":

			im, err := png.Decode(r)
			if err != nil {
				panic("Bad png")
			}
			f, err := os.OpenFile(newFile, os.O_WRONLY|os.O_CREATE, 0777)
			if err != nil {
				panic("Cannot open file")
			}
			// SET THE EXPECTED SIZE, REDUCE TO 13KB
			dst := image.NewRGBA(image.Rect(200, 200, im.Bounds().Max.X/2, im.Bounds().Max.Y/2))

			// RESIZE
			draw.NearestNeighbor.Scale(dst, dst.Rect, im, im.Bounds(), draw.Over, nil)

			// ENCODE TO OUTPUT
			png.Encode(f, dst)

		case "jpeg":

			im, err := jpeg.Decode(r)
			if err != nil {
				panic("Bad jpeg")
			}

			f, err := os.OpenFile(newFile, os.O_WRONLY|os.O_CREATE, 0777)
			if err != nil {
				panic("Cannot open file")
			}

			// SET THE EXPECTED SIZE, REDUCE TO 13KB
			dst := image.NewRGBA(image.Rect(200, 200, im.Bounds().Max.X/2, im.Bounds().Max.Y/2))

			// RESIZE
			draw.NearestNeighbor.Scale(dst, dst.Rect, im, im.Bounds(), draw.Over, nil)

			// ENCODE TO OUTPUT
			jpeg.Encode(f, dst, nil)

		case "gif":

			im, err := gif.Decode(r)
			if err != nil {
				panic("Bad gif")
			}

			f, err := os.OpenFile(newFile, os.O_WRONLY|os.O_CREATE, 0777)
			if err != nil {
				panic("Cannot open file")
			}
			// SET THE EXPECTED SIZE, REDUCE TO 13KB
			dst := image.NewRGBA(image.Rect(200, 200, im.Bounds().Max.X/2, im.Bounds().Max.Y/2))

			// RESIZE
			draw.NearestNeighbor.Scale(dst, dst.Rect, im, im.Bounds(), draw.Over, nil)

			//ENCODE TO OUTPUT
			gif.Encode(f, dst, nil)

		}
	} else {
		newFile = "pix.png"
	}

	if user.Userpicture != "" {
		// url = "https://golang-api-10.herokuapp.com/" + newFile[2:]
		var url = "http://localhost:9000/" + newFile[2:]
		users.Userpicture = url
	}

	users.Lastname = user.Lastname
	users.Firstname = user.Firstname
	users.Email = user.Email
	users.Mobile = user.Mobile
	users.Isactivated = user.Isactivated
	users.Updated_at = now2
	_, err4 := db.Model(users).Where("id = ?", idno).UpdateNotZero()
	if err4 != nil {
		c.JSON(404, err4.Error())
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "User ID has been updated."})
	}

}
