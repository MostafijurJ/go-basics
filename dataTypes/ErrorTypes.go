package dataTypes

import (
	"errors"
	"fmt"
)

func ErrorHandle(){
	fmt.Println("----------- Calling from Error types --------------")
	err := check(0,0)
	call := check(10,0)
	
	fmt.Println(err)

	if call != nil {
		fmt.Println("No error found in call method")
	}

	
}

func check(a,b int) error {
	if a == 0 && b == 0 {
		fmt.Println("In error occured Methods")
		return errors.New("both values are zero, try with another")
	}

	fmt.Println("No error found")
	return nil
}