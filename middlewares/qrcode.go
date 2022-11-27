package middlewares

import (
	"bytes"
	"fmt"
	"image/png"
	"io/ioutil"

	"src/flutter-api/config"
	"src/flutter-api/models"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func Qrcode(idno string) {

	db := config.Connect()
	defer db.Close()

	var tmpUser = new(models.Mfa)
	if err2 := db.Model(tmpUser).Where("id = ?", idno).Select(); err2 != nil {
		fmt.Println(err2.Error())
		return
	} else {
		// GENERATE TIME BASED TOKEN
		key, err := totp.Generate(totp.GenerateOpts{
			Issuer:      tmpUser.Username,
			AccountName: tmpUser.Email,
		})
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		newFile := "./assets/qrcodes/00" + idno + ".png"

		// Convert TOTP key into a PNG
		var buf bytes.Buffer
		img, err := key.Image(200, 200)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		png.Encode(&buf, img)
		display(key, buf.Bytes(), newFile)
		qrimage := "http://localhost:8082/" + newFile[2:]
		// qrimage := "https://golang-api-10.herokuapp.com/" + newFile[2:]
		sql4 := `UPDATE USERS SET qrcode=$1,secretkey=$2 WHERE id=$3`
		db.Exec(sql4, &qrimage, key.Secret(), &idno)

	}
}

func display(key *otp.Key, data []byte, xfile string) {
	ioutil.WriteFile(xfile, data, 0644)
}
