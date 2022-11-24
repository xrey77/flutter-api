// package middlewaresx

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"

// 	"src/flutter-api/config"
// 	"src/flutter-api/models"

// 	"github.com/gin-gonic/gin"
// )

// func GetProducts(c *gin.Context) {

// 	// _, err := ExtractTokenMetadata(c.Request)
// 	// if err != nil {
// 	// 	c.JSON(http.StatusUnauthorized, "unauthorized")
// 	// 	return
// 	// }

// 	//PAGE NUMBER
// 	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

// 	//PER PAGE
// 	perPage, _ := strconv.Atoi(c.DefaultQuery("perPage", "12"))

// 	// var total_page float64 = math.Ceil(float64(total_rec / int64(perPage)))

// 	var xproducts []models.Sale
// 	db := config.Connect()
// 	defer db.Close()
// 	// ctx, cancel := context.WithTimeout(context.Background(), 55*time.Second)
// 	// defer cancel()

// 	//SEARCH KEY
// 	search := c.Query("find")
// 	//TOTAL RECORDS
// 	var total_rec int64 = totalRecs(search)
// 	totalPages := (total_rec / int64(perPage))
// 	if totalRecs(search)%int64(perPage) > 0 {
// 		totalPages++
// 	}

// 	//QUERY ALL RECORDS

// 	//IF SEARCH KEY IS NOT EMPTY
// 	if search != "" {
// 		var sql = "SELECT id,prod_pic,prod_desc,prod_saleprice FROM products WHERE LOWER(prod_desc) LIKE $1  ORDER BY id"
// 		sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, perPage, (page-1)*perPage)
// 		rows, err := db.Query(sql, "%"+search+"%")
// 		if err != nil {
// 			c.JSON(200, gin.H{"message": err})
// 			return
// 		}
// 		defer rows.Close()
// 		for rows.Next() {
// 			var prods models.Sale
// 			// rows.Scan(&prods.ID, &prods.Prod_name, &prods.Prod_desc, &prods.Prod_stockqty, &prods.Prod_unit, &prods.Prod_cost, &prods.Prod_sell, &prods.Prod_pic, &prods.Prod_category, &prods.Prod_saleprice)
// 			rows.Scan(&prods.ID, &prods.Prod_pic, &prods.Prod_desc, &prods.Prod_saleprice)
// 			xproducts = append(xproducts, prods)
// 		}

// 		// totalPages := (int(total_rec) % perPage) == 0 ? totalPages_pre : totalPages_pre + 1

// 		//RETURN JSON
// 		c.JSON(http.StatusOK, gin.H{"data": xproducts, "totalRecs": total_rec, "page": page, "perPage": perPage, "totalPages": totalPages})

// 		// sql = "SELECT id,prod_name,prod_desc,prod_stockqty,prod_unit,prod_cost,prod_sell," +
// 		// "prod_pic,prod_category,prod_saleprice FROM products WHERE prod_desc LIKE '%%" + search + "%%' ORDER BY id"
// 	} else {
// 		var sql = "SELECT id,prod_pic,prod_desc,prod_saleprice FROM products WHERE prod_pic <> 'null' ORDER BY id"
// 		sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, perPage, (page-1)*perPage)
// 		rows, err := db.Query(sql)
// 		if err != nil {
// 			c.JSON(200, gin.H{"message": err})
// 			return
// 		}
// 		defer rows.Close()
// 		for rows.Next() {
// 			var prods models.Sale
// 			// rows.Scan(&prods.ID, &prods.Prod_name, &prods.Prod_desc, &prods.Prod_stockqty, &prods.Prod_unit, &prods.Prod_cost, &prods.Prod_sell, &prods.Prod_pic, &prods.Prod_category, &prods.Prod_saleprice)
// 			rows.Scan(&prods.ID, &prods.Prod_pic, &prods.Prod_desc, &prods.Prod_saleprice)
// 			xproducts = append(xproducts, prods)
// 		}
// 		c.JSON(http.StatusOK, gin.H{"data": xproducts, "totalRecs": total_rec, "page": page, "perPage": perPage, "totalPages": totalPages})
// 	}

// 	// fmt.Println(sql)
// 	//PAGINATION IMPLEMENTATION
// 	// rows, err := db.QueryContext(ctx, sql)

// 	// rows, err := db.Query(sql, "%"+search+"%")

// }

// func totalRecs(search string) int64 {
// 	db := config.Connect()
// 	defer db.Close()
// 	var count int64
// 	if search == "" {
// 		sql := `SELECT count(*) FROM products WHERE prod_pic <> 'null';`
// 		rowTots := db.QueryRow(sql)
// 		err1 := rowTots.Scan(&count)
// 		if err1 != nil {
// 			return 0
// 		}
// 	} else {
// 		sqlCount := `SELECT count(*) FROM products WHERE LOWER(prod_desc) LIKE $1;`
// 		rowTots := db.QueryRow(sqlCount, "%"+search+"%")
// 		err1 := rowTots.Scan(&count)
// 		if err1 != nil {
// 			return 0
// 		}
// 	}
// 	return count
// }
