package maps

import "fmt"

/*
* Key Type Requirements
Must be comparable (support == and !=)
Cannot be slices, maps, or functions
Structs and arrays are allowed if their fields are comparable
*/
type Point struct {
	X, Y int
}

func MapsKeyValidation() {
	points := map[Point]string{
		{X: 0, Y: 0}: "Origin",
		{X: 1, Y: 2}: "Top Right",
	}

	p := Point{X: 1, Y: 2}
	fmt.Println("Value:", points[p])
}

func ZeroValues() {
	mp := map[string]int{}

	mp["a"] = 1
	mp["b"] = 2

	fmt.Println("Map is : ", mp)

	key := "a"

	value, ok := mp[key]
	if ok {
		fmt.Println("Value found : ", value)

	} else {
		fmt.Println("Key found in the map, Map is: ", mp, ", Key: ", key)
	}

}
