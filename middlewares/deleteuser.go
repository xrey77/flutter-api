package middlewares

import (
	"fmt"
	"net/http"

	"src/flutter-api/config"
	"src/flutter-api/models"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized Access")
		return
	}
	idno := c.Param("id")
	db := config.Connect()
	defer db.Close()
	var usermodel = new(models.Users)

	_, err = db.Model(usermodel).Where("id = ?", idno).Delete()
	if err != nil {
		fmt.Println("err 1...." + err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "User ID,  has been deleted."})
	}
}
