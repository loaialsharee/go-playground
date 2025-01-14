// This code snippet is derived from this blog: https://medium.com/@wambuirebeka/advanced-golang-concepts-channels-context-and-interfaces-dc3b71cd0ed8

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(10)
}

func main() {
	dataChan := make(chan int)

	go func() { // go-routine that's running on the background thread
		wg := sync.WaitGroup{}

		for i := 0; i < 10; i++ {
			wg.Add(1)

			go func() { // go-routine that's running on the background thread
				defer wg.Done()
				result := DoWork()
				dataChan <- result
			}()
		}

		wg.Wait()
		close(dataChan)
	}()

	for x := range dataChan {
		fmt.Printf("%d\n", x)

	}

}
