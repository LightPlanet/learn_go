package main

import "fmt"

// Amimal

type Animal interface {
	Speak() string
}

func SaySomething(this Animal) string {
	return this.Speak()
}

// Cat

type Cat struct {
	Name  string
	Voice string
}

func (this Cat) Speak() string {
	return fmt.Sprintf("Кошка %s мяукает %s", this.Name, this.Voice)
}

// Dog

type Dog struct {
	Name  string
	Voice string
}

func (this Dog) Speak() string {
	return fmt.Sprintf("Собака %s гавкает %s", this.Name, this.Voice)
}

func main() {
	cat := Cat{
		Name:  "Матрёшка",
		Voice: "МЯЯУУУУ!",
	}
	dog := Dog{
		Name:  "Кусака",
		Voice: "Гав! Гав! Гав!",
	}
	fmt.Println(SaySomething(cat), SaySomething(dog))
}
