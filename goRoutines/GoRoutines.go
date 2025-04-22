package goRoutines

import (
	"fmt"
	"log"
	"time"
)

const FixedValue = 5

func RunFunctions() {
	runGeneralMethod()
	runGoRoutinesMethods()
}

func runGoRoutinesMethods() {
	fmt.Println("--------------  In Go Routines Start ------------------- ")
	for i := 1; i < 6; i++ {
		go printAsync(i, FixedValue) // create separate thread for each call
	}

	time.Sleep(1 * time.Second)

	fmt.Println(" ------------------------- Go Routine End -------------------- ")
}

func runGeneralMethod() {
	fmt.Println("------------------- In general run method  Start ---------------------")
	for i := 1; i < 6; i++ {
		printAsync(i, FixedValue)
	}

	log.Printf("Hello loggging")
	fmt.Println(" ------------------------- General Run End ----------------------- ")
}

func printAsync(start, finish int) {
	for i := start; i <= finish; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	time.Sleep(1000 * time.Microsecond)
}
