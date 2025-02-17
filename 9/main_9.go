package main

import "fmt"

func cube(num uint8) float64 {
	f := float64(num)
	return f * f * f
}

func cubePipeline(in <-chan uint8) <-chan float64 {
	out := make(chan float64)

	go func() {
		defer close(out)
		for num := range in {
			out <- cube(num)
		}
	}()

	return out
}

func main() {
	inputCh := make(chan uint8)

	outputCh := cubePipeline(inputCh)

	go func() {
		for _, num := range []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9} {
			inputCh <- num
		}
		close(inputCh)
	}()

	for result := range outputCh {
		fmt.Println(result)
	}
}
