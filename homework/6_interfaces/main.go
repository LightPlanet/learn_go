package main

import (
	"fmt"
	"math"
	"time"
)

// Задача 1 -----------------------------------------------------------------------------
// Описание:
// Создайте интерфейс PaymentProcessor с методом Process(amount float64) string,
// который возвращает строку с результатом обработки платежа.
// Реализуйте два типа платежных систем: CreditCard и CryptoWallet,
// каждый из которых реализует интерфейс PaymentProcessor.
// В функции main создайте список платежных систем и вызовите метод Process для каждого,
// выводя результат на экран.
// Требования:
//     Интерфейс PaymentProcessor.
//     Структуры CreditCard и CryptoWallet, реализующие интерфейс.
//     В main создайте массив/слайс этих структур и вызовите их методы.
//     P.S. Подумайте, какие поля могут быть у каждой структуры

type PaymentProcessor interface {
	Process(amount float64) string
}

type CreditCard struct {
	Number         int64
	Owner          string
	ExpirationTime time.Time
}

func (this CreditCard) Process(amount float64) string {
	return fmt.Sprintf(
		"[%d|%s|%s] process amount: %f",
		this.Number, this.Owner, this.ExpirationTime, amount,
	)
}

type CryptoWallet struct {
	Number []byte
	// PublicKey  []byte // Need to send public key as Process argument?
	// PrivateKey []byte
}

func (this CryptoWallet) Process(amount float64) string {
	return fmt.Sprintf(
		"[%s] process amount: %f",
		this.Number, amount,
	)
}

func Ex1() {
	payments := []PaymentProcessor{
		CreditCard{
			1234567812345678,
			"Ivan Ivanov",
			time.Date(2030, time.December, 10, 0, 0, 0, 0, time.UTC),
		},
		CryptoWallet{
			[]byte("bc1q123456789012345678901234567890"),
			// []byte("1234567890123456789012345678901234567890"),
			// []byte("1234567890123456789012345678901234567890"),
		},
	}
	for i, p := range payments {
		fmt.Printf("Payment #%d: %s\n", i, p.Process(999))
	}
}

// Задача 2 -----------------------------------------------------------------------------
// Описание:
// Создайте интерфейс Shape с методом Area() float64.
// Реализуйте структуры Circle и Rectangle, которые реализуют этот интерфейс.
// Напишите функцию, принимающую слайс фигур ([]Shape) и выводящую площадь каждой фигуры.
// Требования:
//     Интерфейс Shape.
//     Структуры Circle (с радиусом) и Rectangle (с длиной и шириной).
//     Функция для вывода площадей всех фигур в слайсе.

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (this Circle) Area() float64 {
	return 2 * math.Pi * this.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (this Rectangle) Area() float64 {
	return this.Width * this.Height
}

func PrintShapeAreas(shapes []Shape) {
	for i, s := range shapes {
		fmt.Printf("Area #%d: %f\n", i, s.Area())
	}
}

func Ex2() {
	shapes := []Shape{
		Circle{10},
		Rectangle{10, 10},
	}
	PrintShapeAreas(shapes)
}

func main() {
	fmt.Println("Ex1")
	Ex1()
	fmt.Println("\nEx2")
	Ex2()
}
