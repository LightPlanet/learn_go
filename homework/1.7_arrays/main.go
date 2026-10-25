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
	"math/rand"
	"os"
	"slices"
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

// Exercise 1 ---------------------------------------------------------------------------

func Ex1() {
	// Fill
	myArray := [10]int{}
	for i := 0; i < len(myArray); i++ {
		myArray[i] = rand.Intn(100)
	}
	fmt.Println("Generated pseudo-random array:", myArray)

	// Note:
	// 1. 	mySlice := MyArray[:] 					// Not a copy - makes a view
	// 2. 	mySlice := make([]int, len(myArray))	// Alternative
	// 		copy(mySlice, myArray[:])

	// Copy
	mySlice := slices.Clone(myArray[:])
	fmt.Println("Copied to slice:              ", mySlice)

	// Sort (lexicographic order by default)
	slices.Sort(mySlice)
	fmt.Println("Sorted slice:                 ", mySlice)
}

// Exercise 2 ---------------------------------------------------------------------------

const CityNotFound = -1

// Returns the index of the city if found, "CityNotFound" constant otherwise
func FindCity(cities []string, name string) int {
	for i, v := range cities {
		if v == name {
			return i
		}
	}
	return CityNotFound
}

func AddCity(cities *[]string, name string) bool {
	if FindCity(*cities, name) == CityNotFound {
		*cities = append(*cities, name)
		return true
	}
	return false
}

func RemoveCity(cities *[]string, name string) bool {
	where := FindCity(*cities, name)
	if where != CityNotFound {
		*cities = slices.Delete(*cities, where, where+1)
		return true
	}
	return false
}

func Ex2() {
	cities := []string{
		"Moscow",
		"Novosibirsk",
		"Tambov",
		"Sochi",
		"Ryazan",
		"Tula",
	}
	fmt.Println("Cities:      ", cities)

	// Add
	newCity := "Voronezh"
	if AddCity(&cities, newCity) {
		fmt.Println("Added city:  ", cities)
	}

	// Find
	index := FindCity(cities, newCity)
	if index != CityNotFound {
		fmt.Printf("Found city:    %s at: %d\n", cities[index], index)
	}

	// Remove
	if RemoveCity(&cities, newCity) {
		fmt.Println("Removed city:", cities)
	}
}

// Exercise 3 ---------------------------------------------------------------------------

func Ex3() {
	calcWords := func(s string) map[string]int {
		ret := make(map[string]int)
		words := strings.SplitSeq(s, " ") // More efficient than Split
		for w := range words {
			ret[w]++
		}
		return ret
	}

	fmt.Print("Enter the sentence: ")
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
	fmt.Println("Exercise 1")
	Ex1()
	fmt.Println("\nExercise 2")
	Ex2()
	fmt.Println("\nExercise 3")
	Ex3()
}
