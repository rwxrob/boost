package eightball_test

import (
	"fmt"
	"log"

	"github.com/rwxrob/boost/weeks/20/eightball"
)

func ExampleRespond() {
	resp := eightball.Respond()
	log.Println(resp)
	fmt.Println(len(resp) > 0)

	// Output:
	// true
}
