package main

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		slice1   []string
		slice2   []string
		expected []string
	}{
		{[]string{"apple", "banana", "cherry"}, []string{"banana"}, []string{"apple", "cherry"}},
		{[]string{"a", "b", "c"}, []string{"x", "y", "z"}, []string{"a", "b", "c"}},
		{[]string{"one", "two", "three"}, []string{"one", "two", "three"}, []string{}},
	}

	for _, test := range tests {
		result := difference(test.slice1, test.slice2)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Ошибка: ожидалось %v, получено %v", test.expected, result)
		}
	}
}
