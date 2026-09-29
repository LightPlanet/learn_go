package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func SleepGopher(i int, ch chan int) {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second) // sleep 0-10 seconds
	ch <- i
}

func NotBufferedChannelExample() {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go SleepGopher(i, ch)
	}
	timeout := time.After(3 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case id := <-ch:
			fmt.Printf("#%d finished\n", id)
		case <-timeout:
			fmt.Println("TIMEOUT!!")
			return
		}
	}
}

// Pipeline example ---------------------------------------------------------------------

func Source(ch12 chan string) {
	data := []string{
		"hello",
		"my bad",
		"good friends",
		"bad friends",
		"wonderful weather",
		"bad weather",
	}
	for _, s := range data {
		ch12 <- s
	}
	close(ch12)
}

func Analyze(ch12, ch23 chan string) {
	for s := range ch12 {
		if !strings.Contains(s, "bad") {
			ch23 <- s
		}
	}
	close(ch23)
}

func Receive(ch23 chan string) {
	for s := range ch23 {
		fmt.Printf("Received: %s\n", s)
	}
}

func PipelineExample() {
	ch12 := make(chan string)
	ch23 := make(chan string)
	go Source(ch12)
	go Analyze(ch12, ch23)
	Receive(ch23) // not go-routine - blocks current thread
}

// C++ unique_lock ----------------------------------------------------------------------

type MutexLock struct {
	Mutex    *sync.Mutex
	IsLocked bool
}

func MakeMutexLock(mutex *sync.Mutex) *MutexLock {
	mutex.Lock()
	return &MutexLock{mutex, true}
}

func (this *MutexLock) Unlock() {
	if this.IsLocked {
		this.Mutex.Unlock()
		this.IsLocked = false
	}
}

func MutexWrapperExample() {
	mutex := sync.Mutex{}

	lock := MakeMutexLock(&mutex)
	defer lock.Unlock()
	// Some "may panic" logic...
	lock.Unlock()
}

// Wait group ---------------------------------------------------------------------------

func SleepGopher2(i int, wg *sync.WaitGroup) {
	time.Sleep(time.Duration(rand.Intn(7)) * time.Second)
	fmt.Println(i)
	wg.Done()
}

func WaitGroupExample() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go SleepGopher2(i, &wg)
	}
	wg.Wait()
	fmt.Println("Finished")
}

func main() {
	PipelineExample()
}
