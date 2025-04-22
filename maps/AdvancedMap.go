package maps

import "fmt"

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
