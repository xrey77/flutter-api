package middlewares

import (
	"log"
	"time"

	"github.com/xlzd/gotp"
)

func GenerateToken() {
	//GENERATE SECRET
	secretLength := 64
	key := gotp.RandomSecret(secretLength)
	log.Println("secret key", key)

	var otp1 string
	var secret string = "4S62BZNFXXSZLCRO"
	totp := gotp.NewDefaultTOTP(secret)
	totp.Now()

	//GET CURRENT TIMESTAMP
	now := time.Now()
	nsec := now.UnixNano()

	//GENERATE OTP
	otp1 = totp.At(nsec)

	//VERIFY GENERATED TOKEN
	otp2 := totp.Verify(otp1, nsec)
	log.Println("verify", otp2)

	//FOR QRCODE
	acct := totp.ProvisioningUri("GOLANG", "REYNALD")
	log.Println(acct)

}
