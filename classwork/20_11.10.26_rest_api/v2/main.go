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
	{"1", "Кнут", "Искусство программирования", true},
	{"2", "Кнут Грэхем Паташник", "Конкретная математика", true},
	{"3", "Стивен Скиена", "Алгоритмы. Руководство по разработке", true},
	{"4", "Стивен Скиена", "Наука о данных", true},
	{"5", "Владимир Арнольд", "Математичесий тривиум", true},
	{"6", "Владимир Арнольд", "Математические методы классической механики", true},
	{"7", "Валенитина Морозова", "Введение в анализ", true},
	{"8", "Елена Власова", "Ряды", true},
	{"9", "Виноградов", "Основы теории чисел", true},
	{"10", "Сосинский", "Введение в топологию", true},
	{"11", "Успенский", "Простейшие примеры математических доказательств", true},
}

func FilterBooks(filter, author string) (ret []Book) {
	for _, b := range Books {
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
	author := r.URL.Query().Get("author")
	filter := r.URL.Query().Get("filter")

	wantedBooks := FilterBooks(filter, author)

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("tmpl.ParseFiles:", err)
		return
	}
	// Map of any - because contains different types
	err = tmpl.Execute(w, map[string]any{
		"Books":  wantedBooks,
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
