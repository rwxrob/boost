package termcolors_test

import (
	"fmt"

	C "github.com/rwxrob/boost/weeks/19/termcolors"
)

func ExampleBlack() {

	// C "github.com/rwxrob/boost/weeks/19/termcolors"

	fmt.Printf("%q\n", C.Black+"black"+C.Reset)
	fmt.Printf("%q\n", C.Red+"red"+C.Reset)
	fmt.Printf("%q\n", C.Green+"green"+C.Reset)
	fmt.Printf("%q\n", C.Yellow+"yellow"+C.Reset)
	fmt.Printf("%q\n", C.Blue+"blue"+C.Reset)
	fmt.Printf("%q\n", C.Magenta+"magenta"+C.Reset)
	fmt.Printf("%q\n", C.Cyan+"cyan"+C.Reset)
	fmt.Printf("%q\n", C.White+"white"+C.Reset)

	// Output:
	// "\x1b[30mblack\x1b[0m"
	// "\x1b[31mred\x1b[0m"
	// "\x1b[32mgreen\x1b[0m"
	// "\x1b[33myellow\x1b[0m"
	// "\x1b[34mblue\x1b[0m"
	// "\x1b[35mmagenta\x1b[0m"
	// "\x1b[36mcyan\x1b[0m"
	// "\x1b[37mwhite\x1b[0m"
}

func ExampleRand() {
	fmt.Printf("%q", C.Rand())
}
