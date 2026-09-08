package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

func AutoConcatenation() {
	s := `first
		  second
		  third`
	fmt.Println(s)
}

func Typename() {
	fmt.Printf("value: %v typename: %T\n", `hello`, `hello`)
}

func RunesAndUnicode() {
	var exclamation byte = 33 // !
	var pi rune = 960
	var smile rune = 128512 // 😀
	fmt.Printf("%c %c %c\n", exclamation, pi, smile)

	emoji := '😀'
	fmt.Printf("%v\n", emoji) // prints 128512
}

func ProgressBar(width uint, duration time.Duration) {
	if width > math.MaxInt {
		panic("ProgressBar.width id too big")
	}
	_w := int(width)
	step := duration / time.Duration(width)
	for i := range _w {
		s := fmt.Sprintf("%s|%s%s|",
			strings.Repeat("\b", _w+2), // +2x"|"
			strings.Repeat("=", i),
			strings.Repeat(" ", _w-i-1))
		fmt.Print(s)
		time.Sleep(step)
	}
	fmt.Println()
}

func JapaneseAndRuneCount() {
	s := "お前はもう死んでいる  なに？！"
	for i, v := range s {
		fmt.Printf("i: %d, rune: %v\n", i, v)
	}
	fmt.Printf("RuneCountInString: %d\n", utf8.RuneCountInString(s))
}

func Encrypt(s string, step int) (ret string) {
	for _, v := range s {
		ret += string(v + rune(step))
	}
	return
}

func Decrypt(s string, step int) (ret string) {
	for _, v := range s {
		ret += string(v - rune(step))
	}
	return
}

func StringUtils() {
	message := "Hello from lesson 5"
	words := strings.Split(message, " ")
	for _, v := range words {
		fmt.Println(v)
	}

	// strings.Replace()

	fmt.Println(strings.TrimSpace("Hello my friend\n\n\n\n"))
}

func main() {
	fmt.Println(os.Executable())
	fmt.Println(string(os.PathSeparator))

	AutoConcatenation()
	Typename()
	RunesAndUnicode()
	ProgressBar(15, time.Second)
	JapaneseAndRuneCount()

	fmt.Printf("%s\n", Decrypt("jr#wkh#ehvw", 3))
	StringUtils()
}
