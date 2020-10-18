package main

/*
echo server - a server that merely echoes a given response back to a client
*/

import (
	"bufio"
	"io"
	"log"
	"net"
)

// echo is a handler function that simply echoes received data.
func echo(conn net.Conn) {
	// A defer statement defers the execution of a function until the
	// surrounding function returns
	defer conn.Close()

	// Create a buffer to store received data.
	b := make([]byte, 512)
	for {
		// Receive data via conn.Read into a buffer.
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected error")
			break
		}
		log.Printf("Received %d bytes: %s", size, string(b))

		// Send data via conn.Write.
		log.Println("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
}

// the bufio package wraps Reader and Writer to create a buffered I/O mechanism
func echo_v2(conn net.Conn) {
	defer conn.Close()

	// initializing a new buffered Reader and Writer
	reader := bufio.NewReader(conn)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("Unable to read data")
		}
		log.Printf("Read %d bytes: %s", len(s), s)

		log.Println("Writing data")
		writer := bufio.NewWriter(conn)
		// write the string to the socket
		if _, err := writer.WriteString(s); err != nil {
			log.Fatalln("Unable to write data")
		}
		// flush write all the data to the underlying writer
		writer.Flush()
	}
}

// Copy(Writer, Reader) convenience function
func echo_v3(conn net.Conn) {
	defer conn.Close()

	// Copy data from io.Reader to io.Writer via io.Copy().
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}

func main() {
	// Bind to TCP port 2020 on all interfaces.
	listener, err := net.Listen("tcp", ":2020")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:2020")
	// infinite loop ensures that the server will continue to listen for
	// connection
	for {
		// Wait for connection. Create net.Conn on connection established.
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// Handle the connection. Using goroutine for concurrency.
		// With the go keyword, making it a concurrent call so that
		// other connections don't block while waiting for the handler
		// function to complete
		go echo_v2(conn)
	}
}
