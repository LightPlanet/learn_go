package handlers

import (
	"html/template"
	"log"
	"net/http"
	"rest_api/models"
)

func sendBookList(w http.ResponseWriter, books []models.Book) {
	tmpl, err := template.ParseFiles("templates/books.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err = tmpl.Execute(w, books); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func ShowBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/books" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	sendBookList(w, models.Books)
}

func ShowAvailBooks(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/available" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	sendBookList(w, models.GetAvailBooks())
}

func ShowAuthor(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/show-author" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	author := r.FormValue("author")
	log.Println("Got author:", author)
	sendBookList(w, models.FilterByAuthor(author))
}

func ShowTakeBook(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/take-book" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/take_book.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err = tmpl.Execute(w, models.GetAvailBooks()); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func ShowTakeBookSuccess(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/take-book-success" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Update database
	bookID := r.FormValue("book_id")
	// log.Println("Got book_id:", bookID)
	// Is never nil (id is validated in take_book.html)
	takenBook := models.TakeBookByID(bookID)

	// Show page
	tmpl, err := template.ParseFiles("templates/take_book_success.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err = tmpl.Execute(w, *takenBook); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
