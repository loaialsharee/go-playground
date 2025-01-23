package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu      sync.Mutex
	counter int
}

// Increment safely increments the counter.
func (sc *SafeCounter) Increment() {
	sc.mu.Lock()
	defer sc.mu.Unlock() // Unlock the mutex when the function exits
	sc.counter++
}

// Value safely retrieves the current value of the counter.
func (sc *SafeCounter) Value() int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.counter
}

func main() {
	// Initialize the SafeCounter
	counter := SafeCounter{}

	// Use a WaitGroup to wait for all goroutines to complete
	var wg sync.WaitGroup

	// Start 10 goroutines, each incrementing the counter 100 times
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Increment()
				time.Sleep(10 * time.Millisecond) // Simulate some delay
			}
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Safely retrieve and print the final value of the counter
	fmt.Printf("Final Counter Value: %d\n", counter.Value())
}
