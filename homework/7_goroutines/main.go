package main

import (
	"fmt"
	"sync"
)

// Задача 1 -----------------------------------------------------------------------------
// Описание:
// Напишите программу, которая запускает три горутины.
// Каждая горутина должна выводить сообщение о выполнении своей задачи и завершаться.
// Основная функция должна дождаться завершения всех горутин перед завершением программы.

func Ex1() {
	// Naive implementation
	printNumber := func(i int, wg *sync.WaitGroup) {
		fmt.Println("Number:", i)
		wg.Done()
	}
	var wg sync.WaitGroup
	wg.Add(3)
	go printNumber(0, &wg)
	go printNumber(1, &wg)
	go printNumber(2, &wg)
	wg.Wait()

	// Better
	printNumber2 := func(i int) {
		fmt.Println("Number2:", i)
	}
	wg.Go(func() { printNumber2(0) })
	wg.Go(func() { printNumber2(1) })
	wg.Go(func() { printNumber2(2) })
	wg.Wait()
}

// Задача 2 -----------------------------------------------------------------------------
// Описание:
// Создайте программу, в которой одна горутина генерирует числа от 1 до 10 и
// отправляет их через канал.
// Другая горутина читает числа из канала и выводит их на экран.
// После отправки всех чисел закройте канал и завершите программу.
// Требования:
//     Используйте каналы для передачи данных.
//     Горутина-генератор должна отправлять числа и закрывать канал после этого.
//     Горутина-потребитель должна читать из канала и выводить числа, пока канал не закрыт.

// WARNING: closes the channel after generation is finished
func GenerateTo(init int, count int, output *chan int) {
	for i := range count {
		*output <- init + i
	}
	close(*output)
}

func PrintFrom(input *chan int) {
	i := 0
	for v := range *input {
		fmt.Printf("#%d: %d\n", i, v)
		i++
	}
}

func Ex2() {
	integers := make(chan int)
	var wg sync.WaitGroup
	wg.Go(func() { GenerateTo(0, 10, &integers) })
	wg.Go(func() { PrintFrom(&integers) })
	wg.Wait()
}

func main() {
	fmt.Println("Ex1")
	Ex1()
	fmt.Println("\nEx2")
	Ex2()
}
