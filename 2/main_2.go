package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateSlice(size int) []int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	slice := make([]int, size)
	for i := range slice {
		slice[i] = r.Intn(100)
	}
	return slice
}

func sliceExample(input []int) []int {
	var result []int
	for _, num := range input {
		if num%2 == 0 {
			result = append(result, num)
		}
	}
	return result
}

func addElements(input []int, value int) []int {
	return append(input, value)
}

func copySlice(input []int) []int {
	result := make([]int, len(input))
	copy(result, input)
	return result
}

func removeElement(input []int, index int) []int {
	if index < 0 || index >= len(input) {
		return input
	}
	return append(input[:index], input[index+1:]...)
}

func main() {
	originalSlice := generateSlice(10)
	fmt.Println("Исходный слайс:", originalSlice)

	evenSlice := sliceExample(originalSlice)
	fmt.Println("Слайс с четными числами:", evenSlice)

	newSlice := addElements(originalSlice, 77)
	fmt.Println("Слайс после добавления 77:", newSlice)

	copiedSlice := copySlice(originalSlice)
	fmt.Println("Копия слайса:", copiedSlice)

	removedSlice := removeElement(originalSlice, 2)
	fmt.Println("Слайс после удаления элемента с индексом 2:", removedSlice)
}
