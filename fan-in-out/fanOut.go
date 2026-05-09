package main

import (
	"fmt"
	"sync"
)

// work generator
func generator(max int) <-chan int {
	out := make(chan int)
	// create a goroutine
	go func() {
		// close a channel
		defer close(out)
		for i := 1; i <= max; i++ {
			out <- i
		}
	}()
	return out
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		result := job * 2
		fmt.Printf("Worker ID: %d - processed job: %d | result: %d \n", id, job, result)
		results <- result
	}
}

var WORKERS = 3

func processFanOut() {
	// sender channel
	jobs := generator(10)

	var wg sync.WaitGroup

	// receiver buffered channel
	results := make(chan int, 10)

	for i := 1; i <= WORKERS; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	func() {
		fmt.Println("Im fired now")
		wg.Wait()
		close(results)
	}()

	var total int
	for result := range results {
		total += result
	}

	fmt.Printf("Total sum: %d", total)

}
