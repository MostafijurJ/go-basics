package slices

import (
	"fmt"
	"slices"
	"sort"
)

func VariadicFund(prefix string, worlds ...string) {
	newWorlds := make([]string, len(worlds))
	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}

	fmt.Println("Printed in variadic function: ", newWorlds)
}

func SortedSlice() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	sort.StringSlice(planets).Sort()
	fmt.Println(planets)

	// 3 index slices
	dwarfs := make([]string, 0, 10)
	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
	fmt.Println("Slice : ", dwarfs, " len :", len(dwarfs), " Cap: ", cap(dwarfs))

	dwarfs = append(dwarfs, "KAJOL", "Hello", "ddd", "rrr", "eeee")
	fmt.Println("Slice : ", dwarfs, " len :", len(dwarfs), " Cap: ", cap(dwarfs))
}

func SliceLearn() {
	fmt.Println("Calling from SliceLearn()")
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3, 5}

	fmt.Println("---------initial slices --------------")
	fmt.Println("s1:", s1, "len:", len(s1), "cap:", cap(s1))
	fmt.Println("s2:", s2, "len:", len(s2), "cap:", cap(s2))

	s1 = append(s1, 4, 5) // capacity will be increased to double of the existing slice

	tmp := s1

	s1[0] = 100

	fmt.Println(slices.Equal(s1, s2))

	fmt.Println("s1:", s1, "len:", len(s1), "cap:", cap(s1))
	fmt.Println("s2:", s2, "len:", len(s2), "cap:", cap(s2))
	fmt.Println("temp: ", tmp, "len:", len(tmp), "cap:", cap(tmp))
}
