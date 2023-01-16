package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rafli-lutfi/perpustakaan/model"
	"github.com/rafli-lutfi/perpustakaan/utils"
)

func findUserWithTokenID(userID int) int {
	db := utils.GetDBConnection()

	user := model.User{}

	db.Where("id = ?", userID).First(&user)

	return user.ID
}

func Auth(c *gin.Context) {
	tokenString, err := c.Cookie("session_token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{
			Error:       "error unauthorized user",
			Description: "there is no session token",
		})
		return
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("JWT_TOKEN")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["expiry"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{
				Error:       "error unauthorized user",
				Description: "token is expired",
			})
			return
		}

		idTokenInt, _ := strconv.Atoi(claims["user_id"].(string))

		userID := findUserWithTokenID(idTokenInt)
		if userID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{
				Error:       "error unauthorized user",
				Description: "please login first",
			})
			return
		}

		c.Set("id", userID)
		c.Next()

	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{
			Error:       "error unauthorized user",
			Description: "there is no token",
		})
	}
}
