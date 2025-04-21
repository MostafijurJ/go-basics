package maps

import (
	"fmt"
	"math"
)

func PrintBasicMap() {

	temperature := map[string]float64{
		"A": 52.5,
		"B": 52.7,
	}

	tmp := temperature["C"]

	// print basic map
	fmt.Println("Map is :", temperature)
	fmt.Println("Temperature is: ", tmp)

	// check if the key is existing or not
	if val, ok := temperature["X"]; ok {
		fmt.Println("Val is: ", val)
	} else {
		fmt.Println("Val is: nil")
	}

	// modify a map

	temperature["A"] = 90.9
	copyMap := temperature

	fmt.Println("Modified Map is :", temperature)
	fmt.Println("copy map is : ", copyMap)

	// modify the copy map
	copyMap["B"] = 11.11
	//  if you change any key it will don't create any copy if the existing
	//  it can change the values in the reference , that's why original map's values also get changed
	delete(copyMap, "B")

	fmt.Println("Modified Copy Map is : ", copyMap)
	fmt.Println("Original Map is  : ", temperature)
}

func MapCounter() {

	fmt.Println("-------------- IN map Counter --------------")
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	frequency := make(map[float64]int)
	for _, t := range temperatures {
		frequency[t]++
	}

	for t, num := range frequency {
		fmt.Printf("%+.2f occurs %d times\n", t, num)
	}

	fmt.Println("-------------- leaving map Counter --------------")
}

func Grouping() {

	fmt.Println("-------------- Grouping --------------")
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0, 14.5, 15.6,
	}
	groups := make(map[float64][]float64)
	for _, t := range temperatures {
		g := math.Trunc(t/10) * 10
		groups[g] = append(groups[g], t)
	}

	for g, temperatures := range groups {
		fmt.Printf("%v: %v\n", g, temperatures)
	}

	fmt.Println("-------------- leaving Grouping --------------")

}

func InvertMap() {
	fmt.Println("-------------- InvertMap --------------")

	countryCodes := map[string]string{
		"BD": "Bangladesh",
		"US": "United States",
		"IN": "India",
		"JP": "Japan",
	}

	inverted := make(map[string]string)
	for code, country := range countryCodes {
		inverted[country] = code
	}

	fmt.Println("Original Map:", countryCodes)
	fmt.Println("Inverted Map:", inverted)

	fmt.Println("-------------- leaving InvertMap --------------")
}
