// На ЯП написать консольное приложение по такому ТЗ:
// Пишем консольное приложение для Библиотекарей
// {
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	ID     string
	Author string
	Title  string
	Avail  bool
}

func (this Book) Print(rowFormat string) {
	fmt.Printf(rowFormat, this.ID, this.Author, this.Title, this.Avail)
}

func PrintHelp() {

}

func FindBook(books []Book, s string) {

}

func calcWidth(books []Book, getString func(Book) string) (ret int) {
	for _, b := range books {
		l := len(getString(b))
		if l > ret {
			ret = l
		}
	}
	return
}

func main() {
	books := []Book{
		{"0", "Donald Knuth", "The Art of Computer Programming", true},
		{"1", "Steven Skiena ", "Algorithm Design Manual", true},
		{"2", "Vladimir Arnold ", "A Mathematical Trivium", true},
	}

	// Calculate format string to align book entries
	// Example: "%-5v%-30v%-40v%-5v\n"
	tableRowFormat := func() (f string) {
		f += "%-"
		f += strconv.Itoa(4 + calcWidth(books, func(b Book) string { return b.ID }))
		f += "v%-"
		f += strconv.Itoa(4 + calcWidth(books, func(b Book) string { return b.Author }))
		f += "v%-"
		f += strconv.Itoa(4 + calcWidth(books, func(b Book) string { return b.Title }))
		f += "v%-5v\n" // "Avail", "true", "false"
		return
	}()

	// State machine
	const (
		Help = iota // default
		NotAvailBooks
		GetBook
		BooksOfAuthor
	)
	userWants := Help

	// Main loop
	var userInput string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Start new frame - clear Terminal buffer
		fmt.Print("\033[H\033[2J")

		// Print current state and help
		fmt.Printf(tableRowFormat, "ID", "Author", "Title", "Avail")
		for _, b := range books {
			switch userWants {
			case NotAvailBooks:
				if b.Avail {
					continue
				}
			case BooksOfAuthor:
				if !strings.Contains(b.Author, userInput) {
					continue
				}
			}
			b.Print(tableRowFormat)
		}
		if userWants == Help {
			fmt.Println("Available commands:")
			PrintHelp()
		}

		// User input
		if scanner.Scan() {
			userInput := scanner.Text()
			switch userInput {
			case "0":
				userWants = Help
			case "1":
				userWants = NotAvailBooks
			case "2":
				userWants = GetBook
			case "3":
				userWants = BooksOfAuthor
			}
		}
	}
}
