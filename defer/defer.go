package _defer

import "fmt"

func DeferExample() {

	showDefer()
}

func showDefer() {
	res := sumNamedReturnedValueFunction(10, 20)
	fmt.Println("Main Sum : ", res)

	res2 := defWithoutNamed(10, 20)
	fmt.Println("Main Sum without named return : ", res2)
}

// this way of writing a function is called named return value.
/*
	1. Named return values are declared in the function signature.
	2. They allow you to return a value without explicitly using the return statement.
	3. Named return values can be modified within the function, and their final value is returned when the function completes.
*/
func sumNamedReturnedValueFunction(a, b int) (result int) {
	result = a + b

	showResult := func() {
		result = result + 10
		fmt.Println("Result in defer : ", result)
	}
	defer showResult()
	return
}

/*
	1. In this case, the function does not have named return values.
	2. The result is calculated and returned explicitly.
	3. The defer function modifies the result before it is returned.
	4. The defer function captures the value of result at the time it is called, not at the time of the return statement.
*/

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
