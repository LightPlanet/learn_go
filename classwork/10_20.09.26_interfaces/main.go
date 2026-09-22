package main

import (
	"fmt"
	"time"
)

// interface

type Speaker interface {
	Speak()
}

type Cat struct {
	voice string //meow
}

func (c Cat) Speak() {
	fmt.Println(c.voice)
}

type Dog struct {
	voice string //wuf
}

func (d Dog) Speak() {
	fmt.Println(d.voice)
}

func DoorCall(animal Speaker) {
	animal.Speak()
}

func InterfaceExample() {
	Barsik := Cat{voice: "Meow!"}
	Sharik := Dog{voice: "Wuf!"}
	DoorCall(Barsik)
	DoorCall(Sharik)
}

// go-routine

func SleepGopher(i int, ch chan int) {
	time.Sleep(3 * time.Second)
	ch <- i
}
func GoExample() {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go SleepGopher(i, ch)
	}
	for i := 0; i < 5; i++ {
		id := <-ch
		fmt.Printf("#%d finished\n", id)
	}
	fmt.Println("main finished...")
}

func main() {
	InterfaceExample()
	GoExample()
}
