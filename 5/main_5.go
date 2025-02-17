package main

import "fmt"

func intersection(slice1, slice2 []int) (bool, []int) {
	set := make(map[int]struct{})
	var result []int

	for _, v := range slice2 {
		set[v] = struct{}{}
	}

	for _, v := range slice1 {
		if _, exists := set[v]; exists {
			result = append(result, v)
		}
	}

	return len(result) > 0, result
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	hasIntersection, intersected := intersection(a, b)
	fmt.Println("Есть ли пересечение?", hasIntersection)
	fmt.Println("Пересеченные значения:", intersected)
}
