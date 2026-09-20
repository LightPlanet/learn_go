// 1. Напишите программу, которая запрашивает у пользователя ввод строки,
// 		а затем выводит число - количество символов в строке
// 2. Напишите программу, которая подсчитывает количество гласных букв
// 		(а, е, ё, и, о, у, ы, э, ю, я) в введённой пользователем строке.
// 3. Создайте функцию capitalizeWords(s string) string,
// 		которая преобразует каждое слово в строке так, чтобы первая буква была заглавной, а остальные — строчными.
// 		Например: "привет мир" → "Привет Мир".
// 4. Напишите программу, которая запрашивает у пользователя ввод строки-формулы,
// 		а выводит сообщение о правильности написания круглых скобок, например:
// Пр. 1 вход: (1+1)*(2+2)
// 		выход: Скобки расставлены верно, 2 открывающиеся, 2 закрывающиеся
// Пр. 2 вход: ((1+1) + (2+2) ))
// 		выход: Скобки расставлены неправильно, 3 открывающиеся, 4 закрывающиеся

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// Exercise 1 ---------------------------------------------------------------------------

func Ex1() {
	fmt.Println("Ex1. Введите строку:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := scanner.Text()
		fmt.Println("Кол-во символов: ", len(text))
	}
}

// Exercise 2 ---------------------------------------------------------------------------

func Ex2() {
	isVowel := func(r rune) bool {
		vowels := []rune{'а', 'е', 'ё', 'и', 'о', 'у', 'ы', 'э', 'ю', 'я'}
		return slices.Contains(vowels, r)
	}
	countVowels := func(s string) (count uint) {
		for _, v := range s {
			if isVowel(v) {
				count++
			}
		}
		return
	}

	fmt.Println("Ex2. Введите строку:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := scanner.Text()
		fmt.Println("Кол-во гласных:", countVowels(text))
	}
}

// Exercise 3 ---------------------------------------------------------------------------

func Ex3() {
	capitalizeWords := func(s string) (ret string) {
		previousRuneIsSpace := true
		for _, v := range s {
			if v == ' ' {
				previousRuneIsSpace = true
				ret += " "
			} else if previousRuneIsSpace { // first rune != ' '
				previousRuneIsSpace = false
				ret += strings.ToUpper(string(v))
			} else {
				ret += strings.ToLower(string(v))
			}
		}
		return
	} // capitalizeWords

	fmt.Println("Ex3. Введите строку:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := scanner.Text()
		fmt.Println("Первая заглавная, остальные строчные:", capitalizeWords(text))
	}
}

// Exercise 4 ---------------------------------------------------------------------------

func Ex4() {
	checkBrackets := func(left, right rune, s string) error {
		counter := 0
		for i, v := range s {
			switch v {
			case left:
				counter++
			case right:
				// This check was added after teacher's comment :D
				if counter == 0 {
					return fmt.Errorf("Пара началась с правой скобки, позиция: %d", i)
				}
				counter--
			}
		}
		// Will never be negative because of "started with right" check
		if counter > 0 {
			return fmt.Errorf("Не хватает %d левых скобок", counter)
		}
		return nil
	} // checkBrackets

	fmt.Println("Ex4. Введите формулу:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := scanner.Text()
		err := checkBrackets('(', ')', text)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Скобки расставлены правильно")
		}
	}
}

func main() {
	Ex1()
	Ex2()
	Ex3()
	Ex4()
}
