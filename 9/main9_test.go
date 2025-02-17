package main

import (
	"reflect"
	"testing"
	"time"
)

func TestCubePipeline(t *testing.T) {
	in := make(chan uint8)
	out := cubePipeline(in)

	go func() {
		for _, num := range []uint8{2, 3, 4, 5} {
			in <- num
		}
		close(in)
	}()

	var results []float64

	done := make(chan struct{})
	go func() {
		for result := range out {
			results = append(results, result)
		}
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Error("Таймаут ожидания результатов из канала")
	}

	expected := []float64{8.0, 27.0, 64.0, 125.0}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Ожидалось %v, получено %v", expected, results)
	}
}
