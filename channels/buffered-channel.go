// This code snippet is derived from this blog: https://medium.com/@wambuirebeka/advanced-golang-concepts-channels-context-and-interfaces-dc3b71cd0ed8

package main

import "fmt"

func main() {
	dataChan := make(chan string, 1) // buffered channel

	dataChan <- "Hey Champ!" //add data to the channel
	c := <-dataChan          // get data from the channel

	fmt.Println(c)
}
