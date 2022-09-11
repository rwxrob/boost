package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rwxrob/boost/weeks/17/hello/foo"
)

func main() {
	//println("Hello World!")   // os.Stderr, but not official
	fmt.Println("Hello World!")                           // os.Stdout
	log.Printf("Hello Log World! (PID: %v)", os.Getpid()) // os.Strerr
	foo.Hello()
}
