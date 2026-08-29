// Условие: На вход программе подаются 5 натуральных чисел.
// Задача программы - считать эти числа,
// отсортировать по убыванию и вывести результат сортировки на экран.
// Также вывести на экран другие параметры считанной последовательности в формате:
// "Самое большое число: {число}",
// "Самое маленькое число: {число}",
// "Среднее арифметическое: {число}"

package main

import (
	"bufio"
	"fmt"
	"os"
)

func InputNumber(reader *bufio.Reader, out *uint) bool {
	fmt.Print("Enter the number: ")
	_, err := fmt.Fscan(reader, out)
	if err != nil {
		fmt.Println(err)
		_, _ = reader.ReadString('\n')
		return false
	}
	return true
}

// Сортирует последовательность vals соответственно предикату pred
func Sort(vals []uint, pred func(uint, uint) bool) {
	var lhs *uint
	var rhs *uint
	for i := len(vals) - 1; i > 1; i-- {
		for j := 0; j < i; j++ {
			lhs = &vals[j]
			rhs = &vals[i]
			if !pred(*lhs, *rhs) {
				*lhs, *rhs = *rhs, *lhs
			}
		}
	}
}

func MinMax(vals []uint) (uint, uint, error) {
	if len(vals) == 0 {
		return 0, 0, fmt.Errorf("empty range")
	}
	min := vals[0]
	for i := 1; i < len(vals); i++ {
		if min > vals[i] {
			min = vals[i]
		}
	}
	max := vals[0]
	for i := 1; i < len(vals); i++ {
		if max < vals[i] {
			max = vals[i]
		}
	}
	return min, max, nil
}

func Accumulate(vals []uint) (uint, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("empty range")
	}
	var ret uint // Default-initialized as 0
	for _, v := range vals {
		ret += v
	}
	return ret, nil
}

func main() {
	var vals [5]uint

	// 1. Ввод 5 натуральных чисел
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < len(vals); {
		if InputNumber(reader, &vals[i]) {
			i++
		}
	}

	// 2. Сортировка по убыванию
	greater := func(lhs, rhs uint) bool { return lhs > rhs }
	Sort(vals[:], greater)
	fmt.Println(vals)

	// 3. Поиск наибольшего и наименьшего
	min, max, err := MinMax(vals[:])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Самое маленькое число: %d\nCaмoe большое число: %d\n", min, max)
	}

	// 4. Поиск среднего арифметисческого
	sum, err := Accumulate(vals[:])
	if err != nil {
		fmt.Println(err)
	} else {
		avg := sum / uint(len(vals))
		fmt.Printf("Среднее арифметическое (округлено): %d\n", avg)
	}
}
