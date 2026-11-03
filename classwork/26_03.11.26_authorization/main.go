// go get github.com/golang-jwt/jwt/v5

// curl "http://localhost:8080/users"
// curl "http://localhost:8080/secret"
// curl -X POST "http://localhost:8080/login" -H "Content-Type:application/json" -d '{"email": "Tihon", "password": "12345"}'
// curl "http://localhost:8080/secret" -H "Authorization:eyJhbGciOiJIUzI1Ni..."

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("myserversecretkey") //IN VAULT!

var users = map[string]string{
	"Tihon": "12345",
	"Masha": "qweqwe",
	"Fedya": "Alphapet",
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func getAllUsers(c *gin.Context) {

}

func login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Not valid input structure",
		})
		return
	}
	currentPassword, exists := users[req.Email]
	if !exists || currentPassword != req.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Incorrect Username or Password",
		})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      req.Email,
		"got":        time.Now().Unix(),
		"expired_at": time.Now().Add(2 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Sorry, cannot sign ur token =(",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func checkToken(c *gin.Context) {
	// Authorization: Bearer awdkjawdkjawhdkjawhdakwjdhbj
	var tokenString string
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		tokenString = authHeader
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "U Are not Authorized!",
		})
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Wrong Signature!")
		}
		return secretKey, nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "WrongToken!",
		})
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.JSON(http.StatusOK, gin.H{
			"secret": "This is secret page!!!",
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "U R Not authorized",
		})
	}
}

func main() {
	r := gin.Default()

	r.GET("/users", getAllUsers) //Public
	r.POST("/login", login)      //login handler
	r.GET("/secret", checkToken) //secret handler ONLY WITH TOKEN

	if err := r.Run(":8080"); err != nil {
		fmt.Println(err)
		return
	}
}
