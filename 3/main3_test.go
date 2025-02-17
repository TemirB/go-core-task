package main

import "testing"

func TestStringIntMap(t *testing.T) {
	sim := NewStringIntMap()

	sim.Add("key1", 100)
	if val, ok := sim.Get("key1"); !ok || val != 100 {
		t.Errorf("Ошибка: ожидалось 100, получено %d", val)
	}

	if !sim.Exists("key1") {
		t.Errorf("Ошибка: ключ 'key1' должен существовать")
	}

	sim.Remove("key1")
	if sim.Exists("key1") {
		t.Errorf("Ошибка: ключ 'key1' должен быть удален")
	}

	sim.Add("key2", 200)
	sim.Add("key3", 300)
	copied := sim.Copy()
	copiedMap := copied.data
	if len(copiedMap) != 2 || copiedMap["key2"] != 200 || copiedMap["key3"] != 300 {
		t.Errorf("Ошибка при копировании карты")
	}
}
