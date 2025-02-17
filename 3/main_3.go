package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{data: make(map[string]int)}
}

func (sim *StringIntMap) Add(key string, value int) {
	sim.data[key] = value
}

func (sim *StringIntMap) Remove(key string) {
	delete(sim.data, key)
}

func (sim *StringIntMap) Copy() *StringIntMap {
	newMap := make(map[string]int)
	for k, v := range sim.data {
		newMap[k] = v
	}
	return &StringIntMap{data: newMap}
}

func (sim *StringIntMap) Exists(key string) bool {
	_, exists := sim.data[key]
	return exists
}

func (sim *StringIntMap) Get(key string) (int, bool) {
	value, exists := sim.data[key]
	return value, exists
}

func main() {
	sim := NewStringIntMap()

	sim.Add("one", 1)
	sim.Add("two", 2)
	sim.Add("three", 3)

	fmt.Println("Exists 'two'?", sim.Exists("two"))

	if val, ok := sim.Get("three"); ok {
		fmt.Println("Value for 'three':", val)
	}

	copyMap := sim.Copy()
	fmt.Println("Copied map:", copyMap)

	sim.Remove("two")
	fmt.Println("Exists 'two' after remove?", sim.Exists("two"))
}
