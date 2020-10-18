package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	// a thread-safe way to control concurrency
	// which acts as a synchronized counter
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		// increase an internal counter
		wg.Add(1)
		go func(j int) {
			// decrements the counter by one
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	// Wait() blocks the execution of the goroutines in which it's called
	// will not allow futuer execution until the internal counter reaches zero
	wg.Wait()
}
