package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		// Sprintf(format string, a ...interface{})
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// port is closed or filtered.
			continue
		}
		// Finishing your connection is just polite
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}
