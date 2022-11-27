package middlewares

import (
	"net/http"

	"src/flutter-api/models"

	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"

	"src/flutter-api/config"
)

func ValidateToken(c *gin.Context) {
	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	idno := c.Param("id")
	otpcode := c.Param("otp")
	db := config.Connect()
	defer db.Close()

	var tmpUsers = new(models.Mfaverify)
	if err2x := db.Model(tmpUsers).Where("id = ?", idno).Select(); err2x != nil {
		c.JSON(400, gin.H{"message": err2x.Error()})
		return
	} else {
		valid := totp.Validate(otpcode, tmpUsers.Secretkey)
		if valid {
			c.JSON(200, gin.H{"users": tmpUsers})
		} else {
			c.JSON(400, gin.H{"message": "Invalid otp code."})
		}
	}
}
