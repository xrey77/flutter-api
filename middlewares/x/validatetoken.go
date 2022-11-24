// package middlewaresx

// import (
// 	"net/http"

// 	"src/flutter-api/models"

// 	"github.com/gin-gonic/gin"
// 	"github.com/pquerna/otp/totp"

// 	"src/flutter-api/config"
// )

// func ValidateToken(c *gin.Context) {
// 	_, err := ExtractTokenMetadata(c.Request)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, "unauthorized")
// 		return
// 	}
// 	idno := c.Param("id")
// 	otpcode := c.Param("otp")
// 	db := config.Connect()
// 	defer db.Close()

// 	var tmpUsers models.Mfaverify
// 	sqlFind := `SELECT username,userpicture,secretkey FROM users WHERE id=$1;`
// 	rowUsers := db.QueryRow(sqlFind, &idno)
// 	err2x := rowUsers.Scan(&tmpUsers.Username, &tmpUsers.Userpicture, &tmpUsers.Secretkey)
// 	if err2x != nil {
// 		c.JSON(400, gin.H{"message": err2x.Error()})
// 		return
// 	}
// 	valid := totp.Validate(otpcode, tmpUsers.Secretkey)
// 	if valid {
// 		c.JSON(http.StatusOK, gin.H{"statuscode": "200", "users": tmpUsers})
// 	} else {
// 		c.JSON(http.StatusNotFound, gin.H{"statuscode": "400", "message": "Invalid otp code."})
// 	}
// }
