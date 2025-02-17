package main

import (
	"testing"
)

func TestSliceExample(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	expected := []int{2, 4, 6, 8}
	result := sliceExample(input)

	if len(result) != len(expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}

func TestAddElements(t *testing.T) {
	input := []int{1, 2, 3}
	value := 10
	expected := []int{1, 2, 3, 10}
	result := addElements(input, value)

	if len(result) != len(expected) || result[len(result)-1] != value {
		t.Errorf("Ошибка добавления элемента: %v", result)
	}
}

func TestCopySlice(t *testing.T) {
	input := []int{1, 2, 3}
	result := copySlice(input)

	if &result == &input {
		t.Errorf("Копия слайса не должна ссылаться на оригинал")
	}
}

func TestRemoveElement(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{1, 3, 4}
	result := removeElement(input, 1)

	if len(result) != len(expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, result)
	}
}
