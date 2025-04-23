package dataTypes

import (
	"errors"
	"fmt"
)

func ErrorHandle(){
	fmt.Println("----------- Calling from Error types --------------")
	err := check(0,0)

	if err != nil {
		fmt.Println("Thanks for the Input")
	}else{
		fmt.Println(err)
	}
}

func check(a,b int) error {
	if a == 0 || b == 0 {
		return errors.New("both values are zero, try with another")
	}

	return nil
}