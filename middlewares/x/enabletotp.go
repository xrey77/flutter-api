// package middlewaresx

// import (
// 	"fmt"
// 	"net/http"
// 	"os"

// 	"src/flutter-api/config"
// 	"src/flutter-api/models"

// 	"github.com/gin-gonic/gin"
// )

// func EnableOtp(c *gin.Context) {

// 	_, err := ExtractTokenMetadata(c.Request)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, "unauthorized")
// 		return
// 	}

// 	//GET PARAM ID
// 	idno := c.Param("id")
// 	mode := c.Param("mode")
// 	db := config.Connect()
// 	defer db.Close()

// 	var tmpUser models.Mfa
// 	sqlFindUser := `SELECT id,username,otp,secretkey FROM USERS WHERE id=$1;`
// 	rowUsername := db.QueryRow(sqlFindUser, &idno)
// 	err2 := rowUsername.Scan(&tmpUser.ID, &tmpUser.Username, &tmpUser.Otp, &tmpUser.Secretkey)
// 	if err2 != nil {
// 		fmt.Println(err2)
// 		c.JSON(400, gin.H{"message": err2.Error()})
// 	}

// 	if mode == "0" {
// 		xkey := ""
// 		sql4 := `UPDATE USERS SET otp=$1,secretkey=$2, qrcode=$3 WHERE id=$4`
// 		db.Exec(sql4, &mode, &xkey, "", &idno)
// 		os.Remove("./assets/qrcodes/00" + idno + ".png")

// 		var tmpQrcode models.Qcode
// 		sqlFindUser := `SELECT id,qrcode FROM USERS WHERE id=$1;`
// 		rowUsername := db.QueryRow(sqlFindUser, &idno)
// 		err2 := rowUsername.Scan(&tmpQrcode.ID, &tmpQrcode.Qrcode)
// 		if err2 != nil {
// 			fmt.Println(err2)
// 			c.JSON(400, gin.H{"message": err2.Error()})
// 		}

// 		c.JSON(http.StatusOK, gin.H{"qrcode": tmpQrcode})

// 	} else {
// 		xotp := randomToken()

// 		//GENERATE SECRET
// 		// secretLength := 64
// 		// skey := gotp.RandomSecret(secretLength)

// 		// totp := gotp.NewDefaultTOTP(skey)
// 		// totp.Now()

// 		// accturl := totp.ProvisioningUri("ANGULAR-GOLANG", tmpUser.Username)

// 		Qrcode(idno)

// 		sql5 := `UPDATE USERS SET otp=$1 WHERE id=$2`
// 		db.Exec(sql5, &xotp, &idno)

// 		// var tmpUser models.Mfa
// 		sqlFindUser := `SELECT qrcode FROM USERS WHERE id=$1;`
// 		rowUsername := db.QueryRow(sqlFindUser, &idno)
// 		err2 := rowUsername.Scan(&tmpUser.Qrcode)
// 		if err2 != nil {
// 			fmt.Println(err2)
// 			c.JSON(400, gin.H{"message": err2.Error()})
// 		}

// 		c.JSON(http.StatusOK, gin.H{"qrcode": tmpUser})

// 		// QRCode(tmpUser.Username, "GOLANG-FRAMEWORK", skey, idno)
// 	}

// }

// // https://www.google.com/chart?chs=200x200&chld=M|0&cht=qr&chl=
// // otpauth://totp/Example%3Aalice%40google.com%3Fsecret%3DJBSWY3DPEHPK3PXP%26issuer%3DExample
