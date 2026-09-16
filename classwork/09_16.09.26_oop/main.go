package main

import (
	"encoding/json"
	"fmt"
)

type Celsius float64

type Car struct {
	Name  string `json:"name" xml:"name"`
	Price uint   `json:"price"`
	Info  int    `json:"-"`
}

func jsonSerialization() {
	bytes, err := json.Marshal(Car{
		Name:  "Hyondai",
		Price: 123,
		Info:  42,
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bytes))
	}
}

func main() {
	jsonSerialization()
}
