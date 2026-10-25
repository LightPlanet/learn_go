package main

import (
	"fmt"
)

// Задача 1 -----------------------------------------------------------------------------
// Описание:
// Создайте структуру Book, которая содержит поля:
// 	Title (строка), Author (строка), Year (целое число).
// Добавьте метод GetInfo(), который возвращает строку с информацией о книге в формате:
// 	"Title" by Author, Year.
// Что нужно сделать:
// Объявить структуру Book.
// Реализовать метод GetInfo() для этой структуры.
// Создать экземпляр книги и вывести информацию с помощью метода.

type Book struct {
	Title  string
	Author string
	Year   int
}

func (this Book) GetInfo() string {
	return fmt.Sprintf("Title: %s\nAuthor: %s\nYear: %d", this.Title, this.Author, this.Year)
}

func Ex1() {
	book := Book{
		Author: "Николай Михайлович Карамзин",
		Title:  "История Государства Российского",
		Year:   1818,
	}
	fmt.Println(book.GetInfo())
}

// Задача 2 -----------------------------------------------------------------------------
// Описание:
// Создайте структуру Student, которая содержит поля: имя (Name) и список оценок (Grades []float64).
// Реализуйте метод AverageGrade() float64, который возвращает средний балл студента.
// Что нужно сделать:
// Объявить структуру и метод.
// Создать студента с несколькими оценками и вывести его средний балл.

type Student struct {
	Name   string
	Grades []float64
}

func (this Student) GetAverageGrade() float64 {
	sum := 0.0
	for _, v := range this.Grades {
		sum += v
	}
	return sum / float64(len(this.Grades))
}

func Ex2() {
	student := Student{
		"Иван Иванович Иванов",
		[]float64{5, 4, 5, 3, 5, 5},
	}
	fmt.Println(student.Name, "avg grade:", student.GetAverageGrade())
}

// Задача 3 -----------------------------------------------------------------------------
// Описание:
// Создайте структуру BankAccount с полями: Owner (владелец счета) и Balance (текущий баланс).
// Реализуйте методы:
// Deposit(amount float64), увеличивающий баланс.
// Withdraw(amount float64), уменьшающий баланс, если хватает средств, иначе выводит сообщение о недостатке средств.
// Что нужно сделать:
// Объявить структуру и методы.
// Создать счет, пополнить его, попытаться снять деньги и вывести итоговый баланс.

type Kopek int64
type Ruble int64

func RubleToKopek(count Ruble) Kopek {
	return Kopek(count * 100)
}

func KopekToRuble(count Kopek) Ruble {
	return Ruble(count / 100)
}

type BankAccount struct {
	Owner   string
	Balance Kopek
}

func (this BankAccount) String() string {
	rub := KopekToRuble(this.Balance)
	kop := this.Balance - RubleToKopek(rub)
	return fmt.Sprintf("Owner: %s\tBalance: %d rubles %d kopeks\n", this.Owner, rub, kop)
}

func (this *BankAccount) Deposit(count Kopek) error {
	if count < 0 {
		fmt.Printf("Некорректная операция!")
	}
	this.Balance += count
	return nil
}

func (this *BankAccount) Withdraw(count Kopek) error {
	if this.Balance < count {
		return fmt.Errorf("Недостаточно средств!")
	}
	this.Balance -= count
	return nil
}

func Ex3() {
	acc := BankAccount{
		"Данил Павлыгин",
		RubleToKopek(50),
	}
	fmt.Println(acc)

	fmt.Println("After withdraw 10.50 rub: ")
	if err := acc.Withdraw(RubleToKopek(10) + 50); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(acc)
	}

	fmt.Println("After withdraw 100 rub:")
	if err := acc.Withdraw(RubleToKopek(100)); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(acc)
	}

	fmt.Println("\nAfter deposit 15.75 rub")
	if err := acc.Deposit(RubleToKopek(15) + 75); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(acc)
	}
}

func main() {
	fmt.Println("Ex1")
	Ex1()
	fmt.Println("\nEx2")
	Ex2()
	fmt.Println("\nEx3")
	Ex3()
}
