package main

import (
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	sem := New()

	n := 3

	sem.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			defer sem.Done()
			time.Sleep(1 * time.Second)
		}()
	}

	sem.Wait()

	t.Log("Все горутины завершены.")
}
