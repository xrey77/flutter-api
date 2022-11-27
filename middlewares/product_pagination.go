// *** Author : Reynald Marquez-Gragasin
// *** Email : reynald88@yahoo.com
// *** Program Name : FLUTTER-API
// *** Syntax : http://localhost:9000/getproductpages?page=1;find=women
// ***          http://localhost:9000/getproductpages?page=1

package middlewares

import (
	"fmt"
	"log"
	"strconv"

	"src/flutter-api/config"
	"src/flutter-api/models"

	"github.com/gin-gonic/gin"
)

func GetProductpages(c *gin.Context) {

	// _, err := ExtractTokenMetadata(c.Request)
	// if err != nil {
	// 	c.JSON(403, "Unauthorized Access")
	// 	return
	// }

	//PAGE NUMBER
	// page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	page, _ := strconv.Atoi(c.Query("page"))

	//PER PAGE
	perPage, _ := strconv.Atoi(c.DefaultQuery("perPage", "12"))

	// var total_page float64 = math.Ceil(float64(total_rec / int64(perPage)))

	// ctx, cancel := context.WithTimeout(context.Background(), 55*time.Second)
	// defer cancel()

	//SEARCH KEY
	search := c.Query("find")

	//TOTAL RECORDS
	var total_rec int64 = totalRecs(search)
	//TOTAL PAGES
	totalPages := (total_rec / int64(perPage))
	if totalRecs(search)%int64(perPage) > 0 {
		totalPages++
	}
	db := config.Connect()
	defer db.Close()
	//QUERY ALL RECORDS

	//IF SEARCH KEY IS NOT EMPTY
	if search != "" {

		var products []models.Products
		var sql = `SELECT * FROM products WHERE LOWER(prod_desc) LIKE '%' || ? || '%'  ORDER BY id`
		sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, perPage, (page-1)*perPage)
		_, err := db.Query(&products, sql, search)
		if err != nil {
			c.JSON(404, "No data found.")
			return
		} else {
			c.JSON(200, gin.H{"data": products, "totalRecs": total_rec, "page": page, "perPage": perPage, "totalPages": totalPages})
		}

	} else {
		//IF SEARCH KEY IS EMPTY

		var products []models.Products
		//id,prod_pic,prod_desc,prod_saleprice
		var sql = `SELECT * FROM products ORDER BY id`
		sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, perPage, (page-1)*perPage)
		_, err := db.Query(&products, sql)
		if err != nil {
			c.JSON(404, "No data found.")
			return
		} else {
			c.JSON(200, products)
			// c.JSON(200, gin.H{"data": products, "totalRecs": total_rec, "page": page, "perPage": perPage, "totalPages": totalPages})
		}
	}

}

func totalRecs(search string) int64 {
	db := config.Connect()
	defer db.Close()
	var count int64
	var prods models.Products
	if search == "" {

		err := db.Model(&prods).Last()
		if err != nil {
			log.Print(err.Error())
			return 0
		} else {
			return prods.ID
		}

	} else {

		sql := `SELECT count(id) FROM products WHERE LOWER(prod_desc) LIKE ?`
		_, err1 := db.Query(prods, sql, "%"+search+"%")
		if err1 != nil {
			return 0
		} else {
			return count
		}
	}
}
