package main

import (
	"fmt"
	"os"

	"github.com/rwxrob/boost/weeks/20/args/internal"
)

func main() {
	out := internal.Output(os.Args...)
	fmt.Print(out)
}
