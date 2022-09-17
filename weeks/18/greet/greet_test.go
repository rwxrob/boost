package greet_test

import (
	"github.com/rwxrob/boost/weeks/18/greet"
)

func ExampleGreet_with_Arguments() {
	greet.Hi("Rob")
	// Output:
	// Hi, Rob, nice to meet you!
}
