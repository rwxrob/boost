package main

import (
	"fmt"
	"os"
)

func SayHi(args ...string) {
	name := "there"

	if len(args) > 0 {
		name = args[0]
	}

	//fmt.Println(`Hi ` + name + `!`)

	fmt.Printf("Hi %v!\n", name)
}

func main() {
	SayHi(os.Args[1:]...)
}
