package main

import (
	"fmt"
	"io"
	"os"
)

func OpenAndWrite() {
	handleError := func(err error) {
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	}

	// Open file and defer its Close
	file := func() *os.File {
		file, err := os.Create("file.txt")
		handleError(err)
		return file
	}()
	defer file.Close()

	// WriteString() handles rune -> byte conversion
	_, err := file.WriteString("Hello, World!\n")
	handleError(err)
}

func OpenAndRead() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	data := make([]byte, 64)
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		fmt.Println(string(data[:n]))
	}
}

func OpenAndRead2() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	io.Copy(os.Stdout, file)
}

func main() {
	//OpenAndWrite()
	//OpenAndRead()
	//OpenAndRead2()

	fmt.Println("Normal stdout")

	//oldStdout := os.Stdout
	var err error
	os.Stdout, err = os.OpenFile("out.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("File stdout")
}
