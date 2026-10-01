package main

import (
	"errors"
	"fmt"
)

// Error handling -----------------------------------------------------------------------

var (
	ErrorInF1 error = fmt.Errorf("Error in f1")
	ErrorInF2 error = fmt.Errorf("Error in f2")
	ErrorInF3 error = fmt.Errorf("Error in f3")
	result    string
)

type MyError struct {
	Func string
	Err  error
}

func (e *MyError) Error() string {
	return fmt.Sprintf("error in %s: %v", e.Func, e.Err)
}

func f1() error {
	err := f2()
	if err != nil {
		return &MyError{Func: "f1", Err: fmt.Errorf("%w:%w", ErrorInF1, err)}
	}
	return nil
}

func f2() error {
	err := f3() //nil
	if err != nil {
		return &MyError{Func: "f2", Err: fmt.Errorf("%w:%w", ErrorInF2, err)}
	}
	return &MyError{Func: "f2", Err: ErrorInF2}
}

func f3() error {
	return nil
}

func findErr(err error) {
	for err != nil {
		var mnr *MyError //mnr - My New Error
		if errors.As(err, &mnr) {
			result = mnr.Func
		}
		switch x := err.(type) {
		case interface{ Unwrap() error }: // Technically a C++ dynamic_cast
			err = x.Unwrap()
		case interface{ Unwrap() []error }: //wraped error stores '[]error'
			for _, e := range x.Unwrap() {
				findErr(e)
			}
			return
		default:
			return
		}
	}
}

//func ErrorsExample() {
//	err := f1()
//	if err != nil {
//		findErr(err)
//	}
//	fmt.Printf("In function: %s\n", result)
//}
//
//type MyError struct {
//    Code int
//    Msg  string
//}
//
//func (e *MyError) Error() string { return e.Msg }
//
//func doWork() error {
//    return fmt.Errorf("wrap: %w", &MyError{Code: 404, Msg: "not found"})
//}
//
//func main() {
//    err := doWork()
//    var myErr *MyError
//    if errors.As(err, &myErr) {
//        fmt.Println("Code:", myErr.Code) // Code: 404
//    }
//}

// Composition --------------------------------------------------------------------------

// A single operation of the Pipe[T]
type PipeFn[T any] func(T) (T, error)

func Pipe[T any](value T) func(fns ...PipeFn[T]) (T, error) {
	return func(fns ...PipeFn[T]) (T, error) {
		var err error
		for _, fn := range fns {
			value, err = fn(value)
			if err != nil {
				return value, err
			}
		}
		return value, nil
	}
}

func AddBrackets(left, right rune) PipeFn[string] {
	return func(s string) (string, error) {
		return (string(left) + s + string(right)), nil
	}
}

func PipeExample() {
	s, err := Pipe("Text")(
		AddBrackets('[', ']'),
		AddBrackets('(', ')'),
		AddBrackets('{', '}'),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Modified string: ", s)
}

func main() {
	// ErrorsExample()
	PipeExample()
}
