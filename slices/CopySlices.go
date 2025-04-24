package slices

import (
	"fmt"
	"sort"
)


func CallingSlice(){

	copySlices()
	sliceSorting()

}

func copySlices() {

	a1 := []int{1}
	a2 := []int{-1, -2}
	a5 := []int{10, 11, 12, 13, 14}
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)
	fmt.Println("a5", a5)
	// copy(destination, input)
	// len(a2) > len(a1)
	copy(a1, a2)
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)

}

func sliceSorting() {

	sInts := []int{1, 0, 2, -3, 4, -20}
	sFloats := []float64{1.0, 0.2, 0.22, -3, 4.1, -0.1}
	sStrings := []string{"aa", "a", "A", "Aa", "aab", "AAa"}

	fmt.Println("sInts original:", sInts)

	sort.Ints(sInts)
	fmt.Println("sInts:", sInts)
	
	sort.Sort(sort.Reverse(sort.IntSlice(sInts)))
	fmt.Println("Reverse:", sInts)

	fmt.Println("sFloats original:", sFloats)
	sort.Float64s(sFloats)

	fmt.Println("sFloats:", sFloats)
	sort.Sort(sort.Reverse(sort.Float64Slice(sFloats)))
	fmt.Println("Reverse:", sFloats)


	fmt.Println("sStrings original:", sStrings)
	sort.Strings(sStrings)
	fmt.Println("sStrings:", sStrings)
	
	sort.Sort(sort.Reverse(sort.StringSlice(sStrings)))
	fmt.Println("Reverse:", sStrings)

}
