// package middlewaresx

// import (
// 	"context"
// 	"net/http"
// 	"time"

// 	"src/flutter-api/models"

// 	"src/flutter-api/config"

// 	"github.com/gin-gonic/gin"
// )

// func GetUsers(c *gin.Context) {

// 	_, err := ExtractTokenMetadata(c.Request)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, "unauthorized")
// 		return
// 	}

// 	var xusers []models.TempUsers
// 	db := config.Connect()
// 	defer db.Close()
// 	ctx, cancel := context.WithTimeout(context.Background(), 55*time.Second)
// 	defer cancel()

// 	rows, err := db.QueryContext(ctx, "SELECT id, firstname, lastname, username,mobile,userpicture,email,isactivated FROM users ORDER BY id")
// 	if err != nil {
// 		c.JSON(200, gin.H{"message": err})
// 		return
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var users models.TempUsers
// 		rows.Scan(&users.ID, &users.Firstname, &users.Lastname, &users.Username, &users.Mobile, &users.Userpicture, &users.Email, &users.Isactivated)
// 		xusers = append(xusers, users)
// 	}
// 	c.JSON(200, xusers)

// }
