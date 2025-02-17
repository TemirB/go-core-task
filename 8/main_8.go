package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore struct {
	mu      sync.Mutex
	count   int
	channel chan struct{}
}

func New() *Semaphore {
	return &Semaphore{}
}

func (s *Semaphore) Add(n int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count += n
	s.channel = make(chan struct{}, s.count)
	for i := 0; i < n; i++ {
		s.channel <- struct{}{}
	}
}

func (s *Semaphore) Done() {
	<-s.channel
	s.mu.Lock()
	s.count--
	s.mu.Unlock()
}

func (s *Semaphore) Wait() {
	for {
		s.mu.Lock()
		if s.count == 0 {
			s.mu.Unlock()
			break
		}
		s.mu.Unlock()
	}
}

func main() {
	n := 5
	sem := New()

	sem.Add(n)

	for i := 0; i < n; i++ {
		go func(id int) {
			defer sem.Done()
			fmt.Printf("Горутина %d работает\n", id)
			time.Sleep(time.Millisecond * 300)
			fmt.Printf("Горутина %d закончила\n", id)
		}(i)
	}

	sem.Wait()
	fmt.Println("Все горутины закончили работу")
}
