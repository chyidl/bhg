package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// The struct type to define new data types by specifying the types
// associated fields and methods
type Person struct {
	// Go Uses capitalization to determine scope
	Name string
	Age  int
}

// struct methods
func (p *Person) SayHello() {
	fmt.Println("Hello,", p.Name)
}

type Friend interface {
	// any type that implements the SayHello() methods is a Friend
	SayHello()
}

func Greet(f Friend) {
	f.SayHello()
}

func f() {
	fmt.Println("f function")
}

func strlen(s string, c chan int) {
	// <- indicates whether the data is flowing to or from a channel
	c <- len(s)
}

// custom error define
type MyError string

func (e MyError) Error() string {
	return string(e)
}

// Error Handling
func foo() error {
	return errors.New("Some Error Occurred")
}

// Handling Structured Data
// The most common packages encoding/json, encoding/xml
// They can turn strings to structures, and structures to string
type Foo struct {
	Bar string
	Baz string
}

func main() {
	fmt.Println("Hello, Black Hat Gophers!")

	// Primitive Data Types
	x := int(42)
	y := string("ChyiYaqing")

	fmt.Println(x)

	// Slices and Maps
	// make() to initialize each variable
	var s = make([]string, 0)
	var m = make(map[string]string)
	// append() to add a new item to a slice
	s = append(s, "some string")
	m["some key"] = "some value"
	fmt.Println(s, m)

	// Poiners, Struct, and Interface

	// points to a particular area in memory and allow you to retrieve the
	// value stored there
	var count = int(42)
	ptr := &count
	fmt.Println(*ptr)
	*ptr = 100
	fmt.Println(count)

	var guy = new(Person)
	guy.Name = "Dave"
	Greet(guy)

	// Control Structures
	// Go need brace
	if x == 1 {
		fmt.Println("X is equal to 1")
	} else {
		fmt.Println("X is not equal to 1")
	}

	// Go will execute no more than one matching or default case
	switch y {
	case "foo":
		fmt.Println("Found foo")
	case "bar":
		fmt.Println("Found bar")
	// execute in the none of the other conditions match
	default:
		fmt.Println("Default case")
	}

	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// loop over a collection
	nums := []int{2, 4, 6, 8}
	for idx, val := range nums {
		fmt.Println(idx, val)
	}

	// Concurrency
	// go keyword mean that the program will run function f() concurrently
	go f()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")

	// Go contains a data type called channels that provide a mechanism through
	// which goroutines can synchronize

	// create the integer channel
	c := make(chan int)
	// multiple concurrent call to the strlen
	go strlen("Salutations", c)
	go strlen("World", c)
	// <- operator this time with data flowing from the channel
	// This execution blocks at this line until adequate data can be read from
	// the channel
	a, b := <-c, <-c
	fmt.Println(a, b, a+b)

	// Error handling
	if err := foo(); err != nil {
		fmt.Println("Error Handling")
	}

	// Handling Structured Data
	f := Foo{"Joe Junior", "Hello Shabado"}
	// This Marshal() method encodes the struct to JSON returning a byte slice
	b1, _ := json.Marshal(f)
	fmt.Println(string(b1))
	// decode in via a call to json.Unmarshal(b, &f)
	json.Unmarshal(b1, &f)

}
