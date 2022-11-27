package middlewares

import (
	"net/http"
	"src/flutter-api/config"
	"src/flutter-api/models"

	"github.com/gin-gonic/gin"
)

func DeleteContact(c *gin.Context) {
	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized Access")
		return
	}
	idno := c.Param("id")

	db := config.Connect()
	defer db.Close()

	contactmodel := new(models.Contact)
	_, err = db.Model(contactmodel).Where("id = ?", idno).Delete()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"message": "Contact has been deleted."})
	}
}
