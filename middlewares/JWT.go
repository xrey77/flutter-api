package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"src/flutter-api/utilz"

	"src/flutter-api/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtAuthentication(h http.Handler) gin.HandlerFunc {

	return func(c *gin.Context) {

		h.ServeHTTP(c.Writer, c.Request)
		notAuth := []string{"/user/register", "/user/login"} //List of endpoints that doesn't require auth
		requestPath := c.Request.URL.Path                    //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {

			if value == requestPath {
				h.ServeHTTP(c.Writer, c.Request)
				return
			}
		}

		response := make(map[string]interface{})
		tokenHeader := c.Request.Header.Get("Authorization") //Grab the token from the header

		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			response = utilz.Message(false, "Missing auth token")

			c.AbortWithStatus(403)
			c.Request.Header.Add("Content-Type", "application/json")
			utilz.Respond(c.Writer, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			response = utilz.Message(false, "Invalid/Malformed auth token")
			c.AbortWithStatus(403)
			c.Request.Header.Add("Content-Type", "application/json")
			utilz.Respond(c.Writer, response)
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &models.Token{}
		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("ACCESS_SECRET")), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			response = utilz.Message(false, "Malformed authentication token")
			c.AbortWithStatus(403)
			c.Request.Header.Add("Content-Type", "application/json")
			utilz.Respond(c.Writer, response)
			return
		}
		c.Writer.Status()

		if !token.Valid { //Token is invalid, maybe not signed on this server
			response = utilz.Message(false, "Token is not valid.")
			c.AbortWithStatus(403)
			//c.WriteHeader(http.StatusForbidden)
			c.Request.Header.Add("Content-Type", "application/json")
			utilz.Respond(c.Writer, response)
			return
		}
		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		fmt.Println("User %", tk.UserId) //Useful for monitoring
		ctx := context.WithValue(c.Request.Context(), "user", tk.UserId)
		//req = c.Request.WithContext(ctx)

		h.ServeHTTP(c.Writer, c.Request.WithContext(ctx)) //proceed in the middleware chain!
	}
}
