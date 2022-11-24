package middlewares

import (
	"net/http"

	"src/flutter-api/config"
	"src/flutter-api/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	db := config.Connect()
	defer db.Close()
	var users models.User
	var userx []models.User
	err2 := db.Model(&users).Column("*").Order("id").Select(&userx)
	if err2 != nil {
		c.JSON(200, gin.H{"message": err2.Error()})
		return
	} else {
		c.JSON(200, userx)
	}
}
