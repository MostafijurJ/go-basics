package errorHandling

import (
	"errors"
	"fmt"
	"time"
)

func CallError() {
	fmt.Println("~~~~~~~~~~~ Starting Error Handling ~~~~~~~~~~~")

	//handleDefer()
	//panicInGoRouting()
	//panicInGoRoutingWithoutRecovery()

	customErrorHandling()
	fmt.Println("~~~~~~~~~~~ End of Error Handling ~~~~~~~~~~~")
}

// 🔄 1. panic with Multiple defer Statements (Defer Chain)
func handleDefer() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	defer fmt.Println("defer 3")

	fmt.Println("Before panic")
	panic("Something went wrong")
}

// 🧵 2. panic Inside a Goroutine with Recovery
func panicInGoRouting() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in goroutine:", r)
			}
		}()
		panic("Panic in goroutine")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Main function continues")
}

// 3. Panic inside GoRoutine without Recovery - bad practice
func panicInGoRoutingWithoutRecovery() {
	go func() {
		panic("Panic in goroutine without recovery")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Main function continues without recovery")
}

type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation failed on field '%s': %s", e.Field, e.Msg)
}
func validate(name string) error {
	if name == "" {
		return &ValidationError{Field: "name", Msg: "cannot be empty"}
	}
	return nil
}

func customErrorHandling() {
	err := validate("")
	if err != nil {
		fmt.Println(err)
		var ve *ValidationError
		if errors.As(err, &ve) {
			fmt.Println("Custom error field:", ve.Field)
		}
	}
}
