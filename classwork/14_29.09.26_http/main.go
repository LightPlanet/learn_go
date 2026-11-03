package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

func booksHandler(w http.ResponseWriter, r *http.Request) {
	type User struct {
		Name  string
		Email string
	}

	newUser := User{
		Name:  "Tema",
		Email: "mrcat@example.com",
	}

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/books.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, newUser); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func contactsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Contacts page")
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

func main() {
	http.HandleFunc("/", booksHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contacts", contactsHandler)
	log.Println("Server starting...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
		return
	}
}
