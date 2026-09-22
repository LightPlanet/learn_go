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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Program state (see Context.UserWants variable)
const (
	UserWantsHelp = iota // default
	UserWantsNotAvailBooks
	UserWantsGetBook
	UserWantsConcreteAuthor
)

// Interface helpers --------------------------------------------------------------------

func PrintCurrentMode(userWants int) {
	fmt.Print("Текущий режим: ")
	switch userWants {
	case UserWantsHelp:
		fmt.Println("Помощь")
	case UserWantsNotAvailBooks:
		fmt.Println("Недоступные книги")
	case UserWantsGetBook:
		fmt.Println("Запрос книги")
	case UserWantsConcreteAuthor:
		fmt.Println("Поиск по автору")
	}
}

func PrintHelp() {
	fmt.Println("0 - Помощь (этот список)")
	fmt.Println("1 - Все книги, которые были взяты")
	fmt.Println("2 - Взять книгу по ID")
	fmt.Println("3 - Поиск по автору")
	fmt.Println("4 - Закрыть программу")
}

func ClearTerminal() {
	fmt.Print("\033[H\033[2J")
}

// Context helpers ----------------------------------------------------------------------

type Book struct {
	ID     string
	Author string
	Title  string
	Avail  bool
}

func (this Book) Print(format string) {
	fmt.Printf(format, this.ID, this.Author, this.Title, this.Avail)
}

func FindAndGetAvailBookByID(books []Book, id string) {
	for i := range books {
		b := &books[i]
		if b.Avail && b.ID == id {
			b.Avail = false
		}
	}
}

func CalcWidth(books []Book, getString func(Book) string) (ret int) {
	for _, b := range books {
		l := len(getString(b))
		if l > ret {
			ret = l
		}
	}
	return
}

// Program context ----------------------------------------------------------------------

type Context struct {
	Books     []Book
	UserInput string
	UserWants int
	RowFormat string
}

func MakeContext(books []Book) *Context {
	ret := &Context{}
	ret.Books = books

	// Calculate format string to align book entries
	ret.RowFormat = func() (f string) {
		f += "%-"
		f += strconv.Itoa(4 + CalcWidth(books, func(b Book) string { return b.ID }))
		f += "v%-"
		f += strconv.Itoa(4 + CalcWidth(books, func(b Book) string { return b.Author }))
		f += "v%-"
		f += strconv.Itoa(4 + CalcWidth(books, func(b Book) string { return b.Title }))
		f += "v%-5v\n" // "Доступно", "true", "false"
		return
	}()
	return ret
}

func (this *Context) ParseUserInput(scanner *bufio.Scanner) {
	if !scanner.Scan() {
		return
	}
	this.UserInput = scanner.Text()

	if this.UserWants == UserWantsGetBook {
		FindAndGetAvailBookByID(this.Books, scanner.Text())
		this.UserInput = "0" // fallback to help screen
	}

	switch this.UserInput {
	case "0":
		this.UserWants = UserWantsHelp
	case "1":
		this.UserWants = UserWantsNotAvailBooks
	case "2":
		this.UserWants = UserWantsGetBook
	case "3":
		this.UserWants = UserWantsConcreteAuthor
	case "4":
		ClearTerminal()
		os.Exit(0)
	}
}

func (this *Context) ShowFrame() {
	// Current mode
	PrintCurrentMode(this.UserWants)
	fmt.Println()

	// Book list
	fmt.Printf(this.RowFormat, "ID", "Автор", "Название", "Доступно")
	for _, b := range this.Books {
		switch this.UserWants {
		case UserWantsNotAvailBooks:
			if b.Avail {
				continue
			}
		case UserWantsConcreteAuthor:
			if !strings.Contains(b.Author, this.UserInput) {
				continue
			}
		}
		b.Print(this.RowFormat)
	}

	// Help
	if this.UserWants == UserWantsHelp {
		fmt.Println("\nДоступные команды:")
		PrintHelp()
	}

	// User input
	switch this.UserWants {
	case UserWantsGetBook:
		fmt.Print("\nВведите ID книги: ")
	case UserWantsConcreteAuthor:
		fmt.Print("\nВведите имя автора: ")
	default:
		fmt.Print("\nВведите запрос: ")
	}
}

func main() {
	ctx := MakeContext([]Book{
		{"0", "Donald Knuth", "The Art of Computer Programming", true},
		{"1", "Steven Skiena", "Algorithm Design Manual", true},
		{"2", "Vladimir Arnold", "A Mathematical Trivium", true},
	})

	// Main loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		ClearTerminal()
		ctx.ShowFrame()
		ctx.ParseUserInput(scanner)

		err := scanner.Err()
		if err != nil {
			fmt.Println(err)
			// Make new scanner?
		}
	}
}
