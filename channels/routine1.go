// This code snippet is derived from this blog: https://medium.com/@wambuirebeka/advanced-golang-concepts-channels-context-and-interfaces-dc3b71cd0ed8

// package main

// import "fmt"

// func main() {
// 	dataChan := make(chan string) //adds data to the channel

// 	go func() { // go-routine that's running on the background thread
// 		dataChan <- "Hey Champ!" // gets data from the channel
// 	}()

// 	c := <-dataChan

// 	fmt.Println(c)
// }
