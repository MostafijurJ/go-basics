package main

import (
	"fmt"
	"go-basics/closure"
	"go-basics/problem_solving"
)

func main() {

	pairs := problem_solving.CountPairs([]int{3, 1, 2, 2, 2, 1, 3}, 2)
	pair2 := problem_solving.CountPairs([]int{1, 2, 3, 4}, 1)
	fmt.Println(pairs)
	fmt.Println(pair2)
	closure.Closure()
}
