package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

// Database -----------------------------------------------------------------------------

type Book struct {
	ID     string
	Author string
	Label  string
	Avail  bool
}

var Books = []Book{
	{"1", "Толкин", "Властелин колец", true},
	{"2", "Достоевский", "Преступление и наказание", true},
	{"3", "Толкин", "Хоббит", true},
	{"4", "Достоевский", "Идиот", true},
	{"5", "Солженицын", "Архипелаг ГУЛАГ", true},
}

func FilterBooks(filter, query string) (ret []Book) {
	for _, b := range Books {
		// Case insesitive compare is better for the UX
		if query != "" && !strings.Contains(strings.ToLower(b.Author), strings.ToLower(query)) {
			continue
		}
		switch filter {
		case "avail":
			if !b.Avail {
				continue
			}
		case "not_avail":
			if b.Avail {
				continue
			}
		}
		ret = append(ret, b)
	}
	return
}

func TakeBookByID(id string) { // -> bool ?
	for i := range Books {
		b := &Books[i]
		if b.Avail && b.ID == id {
			b.Avail = false
		}
	}
}

// Handlers -----------------------------------------------------------------------------

func BookListHadler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("author")
	filter := r.URL.Query().Get("filter")

	wantedBooks := FilterBooks(filter, query)

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Map of any - because contains different types
	tmpl.Execute(w, map[string]any{
		"Books":  wantedBooks,
		"Query":  query,
		"Filter": filter,
	})
}

func TakeBookHandler(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("id")
	TakeBookByID(bookID)

	// Go back to main
	http.Redirect(w, r, "/book-list/", http.StatusSeeOther)
}

// Startup ------------------------------------------------------------------------------

func main() {
	http.HandleFunc("/book-list/", BookListHadler)
	http.HandleFunc("/take-book", TakeBookHandler)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
		return
	}
}
