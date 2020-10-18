GO FUNDAMENTALS
===============

* Setting Up a Development Environment
```
1. Downloading and Installing Go 
2. Setting GOROOT to Define the Go Binrary Location 
  /usr/local/go on a *Nix/BSD-based system 
  set GOROOT=/path/to/go
3. Setting GOPATH to Determine the Location of Your Go Workspace 
4. Using Common Go Tool Comands 
  go run: compile and execute the main package 
  go build: compile application and package and dependencies create binary file on disk but doesn't execute your program 
    # reduce the build file size 
    $ go build -ldflags "-w -s"
➜ go build -ldflags "-w -s"

root in bhg/ch-1 at chyidl on  master [!?] via 🐹 v1.15.2
➜ ll
total 3.3M
-rwxr-xr-x 1 root root 1.4M Oct 12 11:11 ch-1
-rwxr-xr-x 1 root root 2.0M Oct 12 10:56 main
-rw-r--r-- 1 root root   91 Oct 12 10:54 main.go
-rw-r--r-- 1 root root  571 Oct 12 11:10 README.m
  
  go doc: interrogate documentation about a package. function, method, or variable.
    - The output that go doc produces is taken directly out of the source code comments.
  ➜ go doc fmt.Println
package fmt // import "fmt"

func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline is
    appended. It returns the number of bytes written and any write error
    encountered.

  go get: to obtain package source code 
  go fmt: automatically formats your source code
    - enforcing the use of proper line breaks, indentation, and brace alignment 
  golint: reports style mistakes 
    - go get -u golang.org/x/lint/golint
``` 

Understanding Go Syntax
======================= 
```
Data Types:
  Primitive types: 
    bool
    string
    int
    int8 
    int16 
    int32 
    int64
    uint 
    uint8
    uint16 
    uint32 
    uint64
    uintptr 
    byte 
    rune
    float32 
    float64
    complex64
    complex128  

Slices and Maps:
  slices: like arrays you can dynamically resize
  maps: unordered lists of key/value pairs 

Pointers, Structs, and Interfaces 
  pointer: points to a particular area in memory and allows you to retrieve the value stored there. 

```

Concurrency
=========== 
```
goroutines: lightweight threds
```
