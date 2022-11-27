package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"src/flutter-api/config"
	"src/flutter-api/models"
	"src/flutter-api/utilz"

	"github.com/gin-gonic/gin"
)

func EnableOtp(c *gin.Context) {

	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	//GET PARAM ID
	idno := c.Param("id")
	mode := c.Param("mode")
	db := config.Connect()
	defer db.Close()

	var tmpUser = new(models.Mfa)
	if err2 := db.Model(tmpUser).Where("id = ?", idno).Select(); err2 != nil {
		fmt.Println(err2)
		c.JSON(400, gin.H{"message": err2.Error()})
	} else {
		if mode == "0" {
			var users = new(models.User)
			mode2, _ := strconv.ParseInt(mode, 10, 64)
			users.Otp = mode2
			users.Secretkey = ""
			users.Qrcode = ""
			_, err4 := db.Model(users).Where("id = ?", idno).UpdateNotZero()
			if err4 != nil {
				log.Print("Unable to decode the request body.", err.Error())
				c.JSON(http.StatusOK, gin.H{"message": "Unable to decode the request body."})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "Record(s) has been updated successfully."})
			}

			os.Remove("./assets/qrcodes/00" + idno + ".png")

			var tmpQrcode = new(models.Qcode)
			if err2 := db.Model(tmpQrcode).Where("id = ?", idno).Select(); err2 != nil {
				fmt.Println(err2)
				c.JSON(400, gin.H{"message": err2.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"qrcode": tmpQrcode})
			}

		} else {
			xotp, _ := strconv.ParseInt(utilz.RandomToken(), 10, 64)

			//GENERATE SECRET
			// secretLength := 64
			// skey := gotp.RandomSecret(secretLength)

			// totp := gotp.NewDefaultTOTP(skey)
			// totp.Now()

			// accturl := totp.ProvisioningUri("ANGULAR-GOLANG", tmpUser.Username)

			Qrcode(idno)

			// sql5 := `UPDATE USERS SET otp=$1 WHERE id=$2`
			// db.Exec(sql5, &xotp, &idno)
			// ENABLE otp

			var user = new(models.User)
			user.Otp = xotp
			_, err4 := db.Model(user).Where("id = ?", idno).UpdateNotZero()
			if err4 != nil {
				log.Print("Unable to decode the request body.", err.Error())
				c.JSON(http.StatusOK, gin.H{"message": "Unable to enable OTP."})
				return
			} else {

				user2 := new(models.User)
				if err2 := db.Model(user2).Where("id = ?", idno).Select(); err2 != nil {
					fmt.Println(err2)
					c.JSON(400, gin.H{"message": err2.Error()})
				}

				c.JSON(http.StatusOK, gin.H{"qrcode": tmpUser})

				// QRCode(tmpUser.Username, "GOLANG-FRAMEWORK", skey, idno)

			}

		}

	}

}

// https://www.google.com/chart?chs=200x200&chld=M|0&cht=qr&chl=
// otpauth://totp/Example%3Aalice%40google.com%3Fsecret%3DJBSWY3DPEHPK3PXP%26issuer%3DExample
