package main

import (
	"log"
	"src/flutter-api/middlewares"
	"text/template"

	// "time"

	// GinPassportFacebook "github.com/durango/gin-passport-facebook"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseGlob("templates/*"))

func init() {
	err1 := godotenv.Load(".env")
	if err1 != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.Default())

	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{
	// 		"http://127.0.0.1", "http://localhost",
	// 		"http://127.0.0.0:5559", "http://127.0.0.1:9100", "http://127.0.0.1:55766",
	// 		"http://127.0.0.1:5554", "http://127.0.0.1:5555", "http://127.0.0.1:5558", "http://localhost:8000",
	// 		"http://localhost:3000", "http://192.168.1.8:3000",
	// 		"http://localhost:8081", "http://192.168.1.8:8081"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
	// 	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))

	router.Static("/assets", "./assets")
	router.GET("/", func(c *gin.Context) {
		tpl.ExecuteTemplate(c.Writer, "index.html", map[string]interface{}{"MESSAGE": "WELCOME"})
	})

	router.POST("/user/login", middlewares.Login)
	router.POST("/user/register", middlewares.Register)
	router.GET("/getuser/:id", middlewares.GetUserbyID)
	router.GET("/getuserid/:id", middlewares.GetUser)
	router.GET("/getallusers", middlewares.GetUsers)
	// router.DELETE("/deleteuser/:id", middlewares.DeleteUser)
	// router.PUT("/updateuser/:id", middlewares.UpdateUser)
	// router.PUT("/updateusermgt/:id", middlewares.UpdateUserMgt)
	// router.PUT("/enabletotp/:id/:mode", middlewares.EnableOtp)
	// router.PUT("/validatetoken/:id/:otp", middlewares.ValidateToken)
	// router.POST("/createcontact", middlewares.AddContact)
	// router.PUT("/updatecontact/:id", middlewares.UpdateContact)
	// router.GET("/getcontact/:id", middlewares.GetContact)
	// router.GET("/getallcontacts", middlewares.GetAllContacts)
	// router.DELETE("/deletecontact/:id", middlewares.DeleteContact)

	// router.GET("/getdeals", middlewares.GetDeals)
	// router.GET("/getproducts", middlewares.GetProducts)

	// // opts := &oauth2.Config{
	// // 	RedirectURL:  "http://localhost:8080/auth/facebook/callback",
	// // 	ClientID:     "1006766449925076",
	// // 	ClientSecret: "41a5777b7abaf08737e150abb6092ec5",
	// // 	Scopes:       []string{"email", "public_profile"},
	// // 	Endpoint:     facebook.Endpoint,
	// }
	// // auth := router.Group("/auth/facebook")
	// // GinPassportFacebook.Routes(opts, auth)

	// // router.GET("/callback", GinPassportFacebook.middlewares(), func(c *gin.Context) {
	// // 	user, err := GinPassportFacebook.GetProfile(c)
	// // 	fmt.Println(err.Error())
	// // 	fmt.Println(user)
	// // 	if user == nil || err != nil {
	// // 		fmt.Println("failed...")
	// // 		c.AbortWithStatus(500)
	// // 		return
	// // 	}
	// // 	fmt.Println("Got it!...")
	// // 	c.String(200, "Got it!")
	// // })

	// http.ListenAndServe(":9000", router)
	if err := router.Run(":9000"); err != nil {
		log.Print("may error....", err)
	} else {
		log.Print("GIN Server is listnening in port : 9000")
	}

}
