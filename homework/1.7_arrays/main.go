// Задача 1: Работа с массивами и слайсами
// Напишите программу, которая создает массив из 10 целых чисел,
// заполняет его случайными значениями от 1 до 100.
// Затем скопируйте этот массив в слайс и отсортируйте его по возрастанию.
// Выведите исходный массив и отсортированный слайс.
//
// Задача 2: Манипуляции со слайсами
// Создайте слайс строк, содержащий названия городов.
// Реализуйте функции для добавления нового города,
// удаления города по имени и поиска города в списке.
// Продемонстрируйте работу этих функций на примере.
//
// Задача 3: Использование мапов для подсчета частот
// Напишите программу, которая читает строку текста и
// подсчитывает количество вхождений каждого слова.
// Используйте мапу (map[string]int) для хранения результатов.
// Выведите полученную статистику.

package main

import (
	"bufio"
	"fmt"
	"iter"
	"maps"
	"os"
	"strconv"
	"strings"
)

// Helper to make format strings for table-like output
// Returns the max len among entries
func calcWidth(entries iter.Seq[string]) (width int) {
	for entry := range entries {
		l := len(entry)
		if width < l {
			width = l
		}
	}
	return
}

func Ex3() {
	calcWords := func(s string) map[string]int {
		ret := make(map[string]int)
		words := strings.SplitSeq(s, " ") // More efficient than Split
		for w := range words {
			ret[w]++
		}
		return ret
	}

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := scanner.Text()
		resultTable := calcWords(text)

		// The second can be either string or int, so %v is used
		format := "%-" + strconv.Itoa(4+calcWidth(maps.Keys(resultTable))) + "s%v\n"

		fmt.Printf(format, "Word", "Count")
		for key, count := range resultTable {
			fmt.Printf(format, key, count)
		}
	}
}

func main() {
	Ex3()
}
