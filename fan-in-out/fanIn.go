package main

import (
	"fmt"
	"sync"
)

func producer(id string, count int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 0; i < count; i++ {
			out <- fmt.Sprintf("%s-%d", id, i)
		}
	}()
	return out
}

func merge(channels ...<-chan string) <-chan string {
	out := make(chan string)

	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(ch <-chan string) {
			defer wg.Done()
			for val := range ch {
				out <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out

}

func processFanIn() {
	prod1 := producer("A", 5)
	prod2 := producer("B", 5)
	prod3 := producer("C", 5)

	// multiple input channels, merged into one output channel.
	merged := merge(prod1, prod2, prod3)

	for val := range merged {
		fmt.Println("Received:", val)
	}
}
