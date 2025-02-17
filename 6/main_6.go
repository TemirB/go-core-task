package main

import (
	"fmt"
	"math/rand"
	"time"
)

func genRandNums(n int) <-chan int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	out := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			out <- r.Intn(100)
		}
		close(out)
	}()

	return out
}

func main() {
	n := 10
	// fmt.Scan(&n)

	for num := range genRandNums(n) {
		fmt.Println(num)
	}
}
