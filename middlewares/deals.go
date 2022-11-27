package middlewares

import (
	"net/http"
	"src/flutter-api/config"
	"src/flutter-api/models"

	"github.com/gin-gonic/gin"
)

func GetDeals(c *gin.Context) {

	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized Access")
		return
	}

	db := config.Connect()
	defer db.Close()

	var deals models.Deals
	var jsondata []models.Deals
	err2 := db.Model(&deals).Column("*").Select(&jsondata)
	if err2 != nil {
		c.JSON(404, gin.H{"message": err2.Error()})
		return
	} else {
		c.JSON(200, jsondata)
	}

}
