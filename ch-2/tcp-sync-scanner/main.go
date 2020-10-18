package main

import (
	"fmt"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup) {
	// the range to continuously receive from the ports channel, looping until
	// the channel is closed
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}

func main() {
	// create a channel by using make()
	// Buffered channels are ideal for maintaining and tracking work for
	// multiple producers and consumers.
	ports := make(chan int, 100)
	var wg sync.WaitGroup
	// The cap built-in function returns the capacity of v, according to its
	// type
	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}
