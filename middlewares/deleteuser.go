package middlewares

import (
	"database/sql"
	"fmt"
	"net/http"

	"src/flutter-api/config"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	_, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	idno := c.Param("id")
	db := config.Connect()
	defer db.Close()

	query := "DELETE FROM USERS WHERE id=$1"
	var rowsAffected int64
	var result sql.Result

	result, err = db.Exec(query, idno)
	if err != nil {
		fmt.Println("err 1...." + err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		fmt.Println("err 2...." + err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	if rowsAffected > 0 {
		// msg := map[string]string{"msg": "User ID, " + idno + " has been deleted."}
		msg := map[string]string{"msg": "User ID,  has been deleted."}
		c.JSON(http.StatusFound, gin.H{"message": msg})
		return
	} else {
		// c.JSON(http.StatusNotFound, gin.H{"message": "User ID, " + idno + " does not exists."})
		c.JSON(http.StatusNotFound, gin.H{"message": "User ID,  does not exists."})
	}

}
