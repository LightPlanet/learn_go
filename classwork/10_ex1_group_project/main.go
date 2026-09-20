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

// Program state (see main.userWants variable)
const (
	UserWantsHelp = iota // default
	UserWantsNotAvailBooks
	UserWantsGetBook
	UserWantsConcreteAuthor
)

func PrintCurrentMode(userWants int) {
	fmt.Print("Текущий режим: ")
	switch userWants {
	case UserWantsHelp:
		fmt.Println("Помощь")
	case UserWantsNotAvailBooks:
		fmt.Println("Не доступные книги")
	case UserWantsGetBook:
		fmt.Println("Запрос книги")
	case UserWantsConcreteAuthor:
		fmt.Println("Поиск по автору")
	}
}

func PrintHelp() {
	fmt.Println("0 - Помощь (этот список)")
	fmt.Println("1 - Все книги, которые были взяты")
	fmt.Println("2 - Взять книгу по индексу")
	fmt.Println("3 - Книги автора")
	fmt.Println("4 - Закрыть программу")
}

//	func FindBookInUse(books []Book) {
//		for _, v := range books {
//			if v.InUse == true {
//				fmt.Println(v.ID, v.Author, v.Title, v.InUse)
//			}
//		}
//	}
//
//func FindBook(books []Book, s string) {
//	for _, v := range books {
//		if v.Author == s {
//			fmt.Println(v.ID, v.Author, v.Title, v.InUse)
//		}
//	}
//}

func FindAndGetAvailBookByID(books []Book, id string) {
	for i := 0; i < len(books); i++ {
		b := &books[i]
		if !b.Avail {
			continue
		}
		if b.ID == id {
			b.Avail = false
		}
	}
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

func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
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
	userWants := UserWantsHelp

	// Main loop
	var userInput string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		ClearTerminal()

		PrintCurrentMode(userWants)
		fmt.Println()

		// Book list
		fmt.Printf(tableRowFormat, "ID", "Автор", "Название", "Доступно")
		for _, b := range books {
			switch userWants {
			case UserWantsNotAvailBooks: // ok
				if b.Avail {
					continue
				}
			case UserWantsConcreteAuthor:
				if !strings.Contains(b.Author, userInput) {
					continue
				}
			}
			b.Print(tableRowFormat)
		}
		if userWants == UserWantsHelp {
			fmt.Println("\nДоступные команды:")
			PrintHelp()
		}

		// User input
		if userWants == UserWantsGetBook {
			fmt.Print("\nВведите ID книги: ")
		} else {
			fmt.Print("\nВведите запрос: ")
		}
		if scanner.Scan() {
			userInput = scanner.Text() // F*ck ":=". And it is not even a compiler warning!

			if userWants == UserWantsGetBook {
				FindAndGetAvailBookByID(books, scanner.Text())
				userInput = "0" // fallback to help screen
			}

			switch userInput {
			case "0":
				userWants = UserWantsHelp
			case "1":
				userWants = UserWantsNotAvailBooks
			case "2":
				userWants = UserWantsGetBook
			case "3":
				userWants = UserWantsConcreteAuthor
			case "4":
				ClearTerminal()
				os.Exit(0)
			}
		}
	} // main loop
}
