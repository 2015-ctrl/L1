package main

import "fmt"

func intersection(set1, set2 map[int]bool) map[int]bool {
	// Определяем, по какому множеству будем итерироваться
	var smallSet, largeSet map[int]bool
	if len(set1) < len(set2) {
		smallSet, largeSet = set1, set2
	} else {
		smallSet, largeSet = set2, set1
	}

	result := make(map[int]bool) // будет содержать результат пересечения
	for elem := range smallSet {
		if largeSet[elem] {
			result[elem] = true
		}
	}

	return result
}

func main() {
	set1 := map[int]bool{1: true, 2: true, 3: true, 4: true}
	set2 := map[int]bool{3: true, 4: true, 5: true, 6: true, 2: true}
	intersectionResult := intersection(set1, set2)
	fmt.Println(intersectionResult)
}
