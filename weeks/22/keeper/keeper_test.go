package keeper_test

import (
	"strings"

	"github.com/rwxrob/boost/weeks/22/keeper"
)

func ExampleRun() {

	// normally pass os.Stdin, this just for testing,
	// note that colors are disabled unless interactive
	// terminal detected

	r := strings.NewReader("Galahad\n")
	keeper.Run(r)

	// Output:
	// Answer these questions three, 'ere the other side ye see.
	//
	// What is your name?
	// > What is your quest?
	// > What is your favorite colour?
	// >
}
