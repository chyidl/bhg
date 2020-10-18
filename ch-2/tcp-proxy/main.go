package main

import (
	"io"
	"log"
	"net"
)

/*
Proxying a TCP Client

forward all traffic received at http://joesproxy.com to http://hoescatcam.website
*/

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "wwww.apple.com.cn:80")
	if err != nil {
		log.Fatalln("Unable to connect to our unreachable host")
	}
	defer dst.Close()

	// Run in goroutine to prevent io.Copy from blocking
	go func() {
		// Copy our source's output to the destination
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	// Copy our destination's output back to our source
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// Listen on lock port 80
	listener, err := net.Listen("tcp", ":2020")
	if err != nil {
		log.Fatalln("unable to bind to port")
	}

	log.Println("Listening on 0.0.0.0:2020")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn)
	}
}
