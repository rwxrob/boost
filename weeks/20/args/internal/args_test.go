package internal_test

import (
	"fmt"

	"github.com/rwxrob/boost/weeks/20/args"
)

func ExampleOutput() {
	fmt.Println(args.Output("./args", "first", "second", "third"))
	// Output:
	// $0 --> "./args"
	// $1 --> "first"
	// $2 --> "second"
	// $3 --> "third"
}

/*
func TestOutput(t *testing.T) {
	out := output("./args", "first", "second", "third")
	if out != testone {
		t.Errorf("\nwant: %q\ngot:  %q\n", testone, "\n"+out)
	}
}
*/
