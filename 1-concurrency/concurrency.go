package concurrency

import "math/rand/v2"

func Generate(numsCh chan int) {
	defer close(numsCh)

	nums := make([]int, 10)
	for i := 0; i < 10; i++ {
		nums[i] = rand.IntN(101)
	}

	for _, v := range nums {
		numsCh <- v
	}
}

func Square(numsCh chan int, resultCh chan int) {
	defer close(resultCh)

	for v := range numsCh {
		resultCh <- v * v
	}
}
