package middlewares

import (
	"fmt"
	"net/http"

	"src/flutter-api/models"

	"src/flutter-api/config"

	"github.com/gin-gonic/gin"
)

func GetUserbyID(c *gin.Context) {

	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized User")
		return
	}

	id := c.Param("id")
	db := config.Connect()
	defer db.Close()

	tmpUser := new(models.Users)
	if err2 := db.Model(tmpUser).Where("id = ?", id).Select(); err2 != nil {
		fmt.Println(err2)
		c.JSON(400, gin.H{"message": err2.Error()})
		return
	} else {
		c.JSON(200, tmpUser)
	}

}

func GetUser(c *gin.Context) {

	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	id := c.Param("id")
	db := config.Connect()
	defer db.Close()

	tmpUser := new(models.Users)
	if err2 := db.Model(tmpUser).Where("id = ?", id).Select(); err2 != nil {
		c.JSON(400, gin.H{"message": err2.Error()})
		return
	} else {
		c.JSON(200, tmpUser)
	}

}
