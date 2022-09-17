package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rwxrob/boost/weeks/18/greet"
	"github.com/rwxrob/boost/weeks/18/greet/internal"
)

func main() {
	var name string
	var err error
	if len(os.Args) > 1 {
		name = strings.Join(os.Args[1:], " ")
	}
	if name == "" {
		fmt.Println("Hello there, what's your name?")
		name, err = internal.ReadLine(os.Stdin)
		if err != nil {
			log.Print(err)
			return
		}
	}
	greet.Hi(name)
}
