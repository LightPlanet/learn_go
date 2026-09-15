// 1) Пользователь-кассир он вводит в консоль по очереди строки (всего 3 строки):
// 		$enter command:молоко 100(цена) 1(количество)
// 2) введенные данные сохраняются в переменные
// 3) после введения третьего товара - консоль спрашивает: есть ли карта постоянного покупателя?
// 		если да- на все скидка 5%
// 4) резульатом является красиво оформленный чек.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Entry struct {
	Name  string
	Price uint64
	Count uint64
}

func (this Entry) TotalPrice() uint64 {
	return this.Price * this.Count
}

// Allows '%s' use custom struct formatting
func (this Entry) String() string {
	return fmt.Sprintf(
		`%s
%d х %d
Сумма: %d`,
		this.Name, this.Price, this.Count, this.TotalPrice())
}

func extractEntry(s string, out_entry *Entry) bool {
	// Split
	tokens := strings.Split(s, " ")
	// fmt.Println("DEBUG:", tokens)
	if len(tokens) < 3 {
		fmt.Println("Неправильный формат строки, посмотрите подсказку!")
		return false
	}

	// Extract name
	out_entry.Name = tokens[0]

	// Extract price
	var err error
	out_entry.Price, err = strconv.ParseUint(tokens[1], 10, 64)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if out_entry.Price == 0 {
		fmt.Println("В этом магазине нет бесплатных товаров!")
		return false
	}

	// Extract count
	out_entry.Count, err = strconv.ParseUint(tokens[2], 10, 64)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func printPriceList(es *[3]Entry, has_card bool, before_discount, after_discount uint64) {
	get_discount := func() uint64 { // printf helper
		if has_card {
			return 5
		}
		return 0
	}
	fmt.Printf(
		`=======================================
              ОАО Golang
=======================================
%s

%s

%s
---------------------------------------
Итого без скидки: %d
Скидка: %d%%
Итог: %d
---------------------------------------
          Спасибо за покупку!
---------------------------------------
`,
		es[0], es[1], es[2],
		before_discount,
		get_discount(),
		after_discount,
	)
}

func main() {
	var entries [3]Entry
	var before_discount uint64

	// Read & parse
	fmt.Println("Введите товар: '<название> <цена> <количество>'")
	scanner := bufio.NewScanner(os.Stdin) // TODO: make scanner.Err() check
	for i := 0; i < len(entries); {
		entry := &entries[i]
		if scanner.Scan() && extractEntry(scanner.Text(), entry) {
			i++
			before_discount += entry.TotalPrice()
		}
	}

	// Confirm discount
	after_discount := before_discount
	has_card := false
	fmt.Println("Есть ли у Вас карта постоянного покупателя? (да/нет) (yes/no)")
	if scanner.Scan() {
		input := scanner.Text()
		has_card = (input == "да") || (input == "yes")
	}
	if has_card {
		after_discount *= 95
		after_discount /= 100
	}

	// Print result
	printPriceList(&entries, has_card, before_discount, after_discount)
}
