package greeting

import "fmt"

var first_name = "Danil"
var second_name = "Pavlygin"

func SayHi() {
	fmt.Printf("Hello, %s %s!\n", first_name, second_name)
}
