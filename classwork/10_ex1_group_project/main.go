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
	InUse  bool
}

func PrintHelp() {

}

func FindBook(books []Book, s string) {

}

func main() {
	books := []Book{}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Start new frame - clear Terminal buffer
		fmt.Print("\033[H\033[2J")

		// Print current state and help
		for _, book := range books {
			// Print state
		}
		PrintHelp()

		if scanner.Scan() {
			text := scanner.Text()
			// Parse command
		}
	}
}
