package middlewares

import (
	"fmt"
	"net/http"
	"src/flutter-api/config"
	"src/flutter-api/models"

	"github.com/gin-gonic/gin"
)

func GetContact(c *gin.Context) {
	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized Access")
		return
	}
	idno := c.Param("id")
	db := config.Connect()
	defer db.Close()

	jsondata := new(models.Contact)
	if err2 := db.Model(jsondata).Where("id = ?", idno).Select(); err2 != nil {
		fmt.Println(err2)
		c.JSON(400, gin.H{"message": err2.Error()})
		return
	} else {
		c.JSON(200, jsondata)
	}
}
