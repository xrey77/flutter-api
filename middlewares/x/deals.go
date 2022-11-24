// package middlewaresx

// import (
// 	"context"
// 	"time"

// 	"src/flutter-api/config"
// 	"src/flutter-api/models"

// 	"github.com/gin-gonic/gin"
// )

// func GetDeals(c *gin.Context) {

// 	// _, err := ExtractTokenMetadata(c.Request)
// 	// if err != nil {
// 	// 	c.JSON(http.StatusUnauthorized, "unauthorized")
// 	// 	return
// 	// }

// 	var xdeals []models.Deals
// 	db := config.Connect()
// 	defer db.Close()
// 	ctx, cancel := context.WithTimeout(context.Background(), 55*time.Second)
// 	defer cancel()

// 	rows, err := db.QueryContext(ctx, "SELECT deals_picture FROM deals ORDER BY id")
// 	if err != nil {
// 		c.JSON(200, gin.H{"message": err})
// 		return
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var deals models.Deals
// 		rows.Scan(&deals.Deals_picture)
// 		xdeals = append(xdeals, deals)
// 	}
// 	c.JSON(200, xdeals)

// }
