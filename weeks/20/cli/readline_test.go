package cli_test

import (
	"fmt"
	"log"
	"strings"

	"github.com/rwxrob/boost/weeks/20/cli"
)

func ExampleReadLine() {
	sr := strings.NewReader("Sample\r\n")
	line, err := cli.ReadLine(sr)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	fmt.Printf("%q", line)
	// Output:
	// "Sample"
}
