package _defer

import "fmt"

func DeferExample() {

	showDefer()
}

func showDefer() {
	res := sum(10, 20)
	fmt.Println("Main Sum : ", res)

	res2 := defWithoutNamed(10, 20)
	fmt.Println("Main Sum without named return : ", res2)
}

// this way of writing a function is called named return value.
func sum(a, b int) (result int) {
	result = a + b

	showResult := func() {
		result = result + 10
		fmt.Println("Result in defer : ", result)
	}
	defer showResult()
	return
}

func defWithoutNamed(a, b int) int {
	result := a + b

	showResult := func() {
		result = result + 10
		fmt.Println("Result in defer: ", result)
	}
	defer showResult()
	return result
}

func def() {
	i := 0

	fmt.Println("First : ", i)

	defer fmt.Println("Second : ", i) // while def() is running, defer will capture the value of i at this point in time,
	// which is 0. and store this value in a stack.
	i++

	fmt.Println("Third : ", i)

	defer fmt.Println("Fourth : ", i) // defer will capture the value of i at this point in time,
	// which is 1. and store this value in a stack.

	return
}
