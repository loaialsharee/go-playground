package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	message string     // Shared resource
	mutex   sync.Mutex // Mutex to protect message
)

func updateMessage(newMessage string, wg *sync.WaitGroup) {
	defer wg.Done() // Mark goroutine as done

	mutex.Lock()         // Lock before modifying the message
	message = newMessage // Update shared variable
	fmt.Println("Updated Message:", message)
	mutex.Unlock() // Unlock after modification

	time.Sleep(time.Millisecond * 100) // Simulate work
}

func main() {
	var wg sync.WaitGroup

	// Launch multiple goroutines to update the message
	wg.Add(2)
	go updateMessage("Hello, World!", &wg)
	go updateMessage("Mutex in Go", &wg)

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Final Message:", message)
}
