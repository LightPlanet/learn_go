package main

import "fmt"

// Задача 1 -----------------------------------------------------------------------------
// Реализуйте функцию divide(a, b float64) (float64, error), которая делит a на b.
// Если b равно нулю, возвращайте ошибку.
// В основной программе вызовите эту функцию и обработайте возможную ошибку.

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return a, fmt.Errorf("Can't divide by zero!")
	}
	return a / b, nil
}

func Ex1() {
	a, b := 12.34, 2.0
	res, err := Divide(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%f/%f=%f\n", a, b, res)
	}

	a, b = 42, 0
	res, err = Divide(a, b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%f/%f=%f\n", a, b, res)
	}
}

// Задача 2 -----------------------------------------------------------------------------
// Функция высшего порядка: передача функции как аргумента
// Создайте функцию applyOperation(a, b int, op func(int, int) int) int,
// которая применяет переданную функцию op к числам a и b.

// Создайте несколько функций-операций: сложение, вычитание, умножение.
// В основной программе вызовите applyOperation с разными операциями и выведите результаты.

func ApplyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func Ex2() {
	sum := func(a, b int) int { return a + b }
	sub := func(a, b int) int { return a - b }
	mul := func(a, b int) int { return a * b }

	a, b := 7, 23
	fmt.Printf("sum(%d, %d) -> %d\n", a, b, ApplyOperation(a, b, sum))
	fmt.Printf("sub(%d, %d) -> %d\n", a, b, ApplyOperation(a, b, sub))
	fmt.Printf("mul(%d, %d) -> %d\n", a, b, ApplyOperation(a, b, mul))
}

// Задача 3 -----------------------------------------------------------------------------
// Напишите программу, которая:
// 1) Создаёт одну анонимную функцию, которая при каждом вызове
// увеличивает внутренний счётчик на 1 и возвращает его значение.
// 2) Вызывает эту функцию несколько раз и выводит результат.
// Пример:
//     // Вызовы функции
//     fmt.Println(counter()) // 1
//     fmt.Println(counter()) // 2
//     fmt.Println(counter()) // 3

func MakeCounter() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}

func Ex3() {
	counter := MakeCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func main() {
	fmt.Println("Ex1")
	Ex1()
	fmt.Println("\nEx2")
	Ex2()
	fmt.Println("\nEx3")
	Ex3()
}
