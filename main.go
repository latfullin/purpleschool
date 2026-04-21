package main

import (
	"fmt"
	concurrency "kilkenny/purpleschool/1-concurrency"
)

func main() {
	numsCh := make(chan int)
	resultCh := make(chan int)

	go concurrency.Generate(numsCh)
	go concurrency.Square(numsCh, resultCh)

	for v := range resultCh {
		fmt.Printf("%d ", v)
	}

}
