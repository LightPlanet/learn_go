// На ЯП написать консольное приложение по такому ТЗ:
// Пишем консольное приложение для Библиотекарей
// Изначально в библиотеке уже есть несколько книг.
// Комады, доступные для набора:
// 0) вывести помощь
// 1) Вывести все книги, которые были взяты
// 2) Взять книгу по ее индексу (нельзя брать уже взятые книги)
// 3) Вывести все книги по автору (вообще все)

package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strings"

	_ "modernc.org/sqlite"
)

// Database -----------------------------------------------------------------------------

var Database *sql.DB

type Book struct {
	ID     string
	Author string
	Label  string
	Avail  bool
}

// Is called by GetBooksFiltered() & GetBookByID().
// Do fresh SELECT because in theory Database can be changed by other client?
// Hovewer, current implementation does not support real time update :D
func DbSelectBooks() (books []Book) {
	rows, err := Database.Query(`SELECT * FROM 'Books';`)
	if err != nil {
		log.Println("SELECT failed:", err)
		return
	}
	defer rows.Close()

	books = []Book{}
	var avail int // Workaround: SQLite does not have a bool
	for rows.Next() {
		b := Book{}
		err := rows.Scan(&b.ID, &b.Author, &b.Label, &avail)
		if err != nil {
			log.Println("rows.Scan failed:", err)
			continue
		}
		if err := rows.Err(); err != nil {
			log.Println("rows.Err failed:", err)
			continue
		}
		b.Avail = (avail == 1)
		books = append(books, b)
	}
	return
}

func GetBooksFiltered(filter, author string) (filtered []Book) {
	books := DbSelectBooks()
	for _, b := range books {
		// Case insesitive compare is better for the UX
		if author != "" && !strings.Contains(strings.ToLower(b.Author), strings.ToLower(author)) {
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
		filtered = append(filtered, b)
	}
	return
}

func TakeBookByID(id string) { // -> bool ?
	books := DbSelectBooks()
	for i := range books {
		b := &books[i]
		if b.Avail && b.ID == id {
			_, err := Database.Exec(`
				UPDATE Books SET Avail = 0
				WHERE ID = (?);`, b.ID)
			if err != nil {
				log.Println("UPDATE failed:", err)
				return
			}
		}
	}
}

// Handlers -----------------------------------------------------------------------------

func BookListHadler(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	filter := r.URL.Query().Get("filter")

	wantedBooks := GetBooksFiltered(filter, author)

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("tmpl.ParseFiles:", err)
		return
	}
	// Map of any - because contains different types
	err = tmpl.Execute(w, map[string]any{
		"books":  wantedBooks,
		"Author": author,
		"Filter": filter,
	})
	if err != nil {
		log.Println("tmpl.Execute:", err)
	}
}

func TakeBookHandler(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("id")
	TakeBookByID(bookID)

	// Go back to main (StatusSeeOther = 303 - redirection code)
	http.Redirect(w, r, "/book-list/", http.StatusSeeOther)
}

// Startup ------------------------------------------------------------------------------

func main() {
	// Load database
	var err error
	if Database, err = DbInit(); err != nil {
		log.Println("DbInit failed with error:", err)
		return
	}
	defer Database.Close()

	// Setup server
	http.HandleFunc("/book-list/", BookListHadler)
	http.HandleFunc("/take-book", TakeBookHandler)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
		return
	}
}
