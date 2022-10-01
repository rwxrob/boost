package main

import (
	"fmt"

	C "github.com/rwxrob/boost/weeks/19/termcolors"
)

func main() {
	fmt.Print(C.CurOff)
	for {
		fmt.Print(C.Rand() + "nyan" + C.Reset)
		//time.Sleep(10 * time.Millisecond)
	}
}
