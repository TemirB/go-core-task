package main

import (
	"fmt"
	"sync"
)

func merge_channels(channels ...<-chan int) <-chan int {
	merged_channel := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}

		wg.Add(len(channels))

		for _, channel := range channels {
			go func(channel <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for id := range channel {
					merged_channel <- id
				}
			}(channel, wg)
		}

		wg.Wait()
		close(merged_channel)
	}()

	return merged_channel
}

func fill(channel chan int, values []int) chan int {
	for _, value := range values {
		channel <- value
	}
	close(channel)
	return channel
}

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go fill(a, []int{1, 2, 3})
	go fill(b, []int{4, 5, 6})
	go fill(c, []int{7, 8, 9})

	for num := range merge_channels(a, b, c) {
		fmt.Println(num)
	}
}
