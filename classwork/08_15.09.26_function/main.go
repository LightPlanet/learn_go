package main

import "fmt"

func mapExample() {
	myMap := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key2"])

	// Iteration
	for k, v := range myMap {
		fmt.Println(k, "->", v)
	}

	// Deletion
	delete(myMap, "key3")
	fmt.Println(myMap)

	// Append
	myMap["key4"] = "value4"
}

func makeTuple(a ...any) []any {
	return a
}

func Speak(animal string) (voice string) { //func(string)(string)
	switch animal {
	case "cat":
		voice = "Meow"
	case "dog":
		voice = "Wuf"
	default:
		voice = "pffff"
	}
	return
}

func PrintVoice(animal string, how func(string) string) {
	fmt.Println(how(animal))
}

func firstClassCitizen() {
	var voice func(string) string
	voice = Speak
	PrintVoice("cat", voice)
}

func deferOperator() {
	{ // scope has no effect
		defer fmt.Println("deferred 1") // called in a stack order
		fmt.Println("not deferred 1")
	}
	fmt.Println("not deferred 2")
	{
		defer fmt.Println("deferred 2")
		fmt.Println("not deferred 3")
	}
	fmt.Println("not deferred 4")
}

// Prints: "Казалось бы, На самом деле"
func closure() {
	f1 := func() string {
		value := "Казалось бы" // "value" declaration
		defer func() { value = "На самом деле" }()
		return value // local variable is copied
	}
	f2 := func() (value string) { // "value" declaration
		value = "Казалось бы"
		defer func() { value = "На самом деле" }()
		return value // return slot is returned (defer mutates it directly)
	}
	fmt.Println("f1: ", f1())
	fmt.Println("f2: ", f2())
}

func main() {
	// firstClassCitizen()
	// deferOperator()
	closure()
}
