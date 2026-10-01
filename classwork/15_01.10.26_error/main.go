package main

import (
	"errors"
	"fmt"
)

// Error handling -----------------------------------------------------------------------

var ErrorInF1 error = fmt.Errorf("Error in f1")
var ErrorInF2 error = fmt.Errorf("Error in f2")
var ErrorInF3 error = fmt.Errorf("Error in f3")

func f1() error {
	err := f2()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrorInF1, err)
	}
	return nil
}
func f2() error {
	err := f3() //nil
	if err != nil {
		return fmt.Errorf("Error in f2: %w", err)
	}
	return ErrorInF2
}
func f3() error {
	return nil
}

func ErrorsExample() {
	err := f1()
	if err != nil {
		if errors.Is(err, ErrorInF1) {
			fmt.Println("YESSS!")
			return
		}
	}
	fmt.Println("ooops")
}

// Composition --------------------------------------------------------------------------

type PipeFn[T any] func(T) (T, error)

func Pipe[T any](s T) func(fns ...PipeFn[T]) (T, error) {
	return func(fns ...PipeFn[T]) (T, error) {
		for _, fn := range fns {
			var err error
			s, err = fn(s)
			if err != nil {
				return s, err
			}
		}
		return s, nil
	}
}

func AddBrackets(left, right string) PipeFn[string] {
	return func(s string) (string, error) {
		return fmt.Sprintf("%s%s%s", left, s, right), nil
	}
}

func PipeExample() {
	s, err := Pipe("Text")(
		AddBrackets("[", "]"),
		AddBrackets("(", ")"),
		AddBrackets("{", "}"),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}

func main() {
	ErrorsExample()
}
