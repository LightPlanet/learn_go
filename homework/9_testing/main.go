// Generate coverage: go test -coverprofile=coverage.out
// Convert to html: go tool cover -html=coverage.out
// OR: go tool cover -html=coverage.out -o coverage.html

package main

import "fmt"

// Задача 1 из темы 8 -------------------------------------------------------------------
// Реализуйте функцию divide(a, b float64) (float64, error), которая делит a на b.
// Если b равно нулю, возвращайте ошибку.
// В основной программе вызовите эту функцию и обработайте возможную ошибку.

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Divide: divisor is zero")
	}
	return a / b, nil
}

// Задача 4 из темы 8 -------------------------------------------------------------------
// Создайте функцию sumAll, которая принимает целые числа и возвращает их сумму.
// Пример использования:
//     fmt.Println(sumAll(1, 2, 3)) // 6
//     fmt.Println(sumAll(10, -2, 4, 7)) // 19

func SumAll(vals []int) (sum int) {
	for _, v := range vals {
		sum += v
	}
	return
}

func main() {
}
