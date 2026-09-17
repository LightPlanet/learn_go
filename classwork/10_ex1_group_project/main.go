package main

import (
	"bufio"
	"fmt"
	"os"
)

type Book struct {
	ID     string
	Author string
	Title  string
}

func PrintState(avail []Book, inUse []Book) {

}

func PrintHelp() {

}

func FindBook(books []Book, s string) {

}

func main() {
	avail := []Book{}
	var inUse []Book

	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Start new frame - clear Terminal buffer
		fmt.Print("\033[H\033[2J")

		// Print current state and help
		PrintState(avail, inUse)
		PrintHelp()

		if scanner.Scan() {
			text := scanner.Text()
			// Parse command
		}
	}
}
