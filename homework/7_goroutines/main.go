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
// Требования:
//     Используйте sync.WaitGroup для синхронизации.
//     Каждая горутина выводит сообщение вида: "Горутина X начала работу" и "Горутина X завершила работу".
//     В основном потоке дождитесь завершения всех горутин.

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

func main() {
	fmt.Println("Ex1")
	Ex1()
}
