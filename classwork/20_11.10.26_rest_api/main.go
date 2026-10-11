// На ЯП написать консольное приложение по такому ТЗ:
// Пишем консольное приложение для Библиотекарей
// type Book struct {
//     Id
//     Author
//     Title
//     Available
// }
// Изначально в библиотеке уже есть несколько книг.
// Комады, доступные для набора:
// 0) вывести помощь
// 1) Вывести все книги, которые были взяты
// 2) Взять книгу по ее индексу (нельзя брать уже взятые книги)
// 3) Вывести все книги по автору (вообще все)

// GET - (показать все книги), создать слайс с доступными книгами, отправить его в http.ResponseWriter
// Выпадающий список для поиска по автору
// books/take-book - Форма (formValue.put()) ВЗЯТА - НЕ ВЗЯТА

package main

import (
	"log"
	"net/http"
	"rest_api/handlers"
)

func main() {
	http.HandleFunc("/books", handlers.ShowBooksHandler)
	http.HandleFunc("/available", handlers.ShowAvailBooks)
	http.HandleFunc("/take-book", handlers.ShowTakeBook)
	http.HandleFunc("/show-author", handlers.ShowAuthor)
	http.HandleFunc("/take-book-success", handlers.ShowTakeBookSuccess)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println(err)
		return
	}
}
