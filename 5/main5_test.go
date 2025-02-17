package main

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		slice1   []int
		slice2   []int
		expected bool
		result   []int
	}{
		{[]int{65, 3, 58, 678, 64}, []int{64, 2, 3, 43}, true, []int{64, 3}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, false, []int{}},
		{[]int{10, 20, 30}, []int{30, 40, 50}, true, []int{30}},
	}

	for _, test := range tests {
		got, result := intersection(test.slice1, test.slice2)
		if got != test.expected {
			t.Errorf("Ошибка: ожидается %v, получено %v", test.expected, got)
		}
		if !reflect.DeepEqual(result, test.result) {
			t.Errorf("Ошибка: ожидается %v, получено %v", test.result, result)
		}
	}
}
