package closure

import "fmt"

const a = 10

var p = 100

func outer() func() {
	money := 100
	age := 30

	fmt.Println("Age:", age)

	show := func() {
		money = money + a + p
		fmt.Println("Money:", money)
	}

	return show
}

func callOuter() {
	closure := outer()
	closure()
	closure()

	closur2 := outer()
	closur2()
	closur2()
}

func Closure() {
	fmt.Println("Hello World from Closure!")

	callOuter()
	fmt.Println("~~~~~~~~~~~ End of Closure ~~~~~~~~~~~")

}

func init() {
	fmt.Println("~~~~~~~~~~ Starting Closure ~~~~~~~~~~~")
}
