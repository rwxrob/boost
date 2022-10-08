package waffles

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	C "github.com/rwxrob/boost/weeks/19/termcolors"
	"github.com/rwxrob/boost/weeks/20/cli"
)

var YES = regexp.MustCompile(`(?i)^(y(es|eah|ep)$|da|affirmative|sure|ok|si)`)

func prompt(in io.Reader) string {
	resp, err := cli.Prompt(in, C.M+"> "+C.Y)
	if err != nil {
		log.Println(err)
		return ""
	}
	fmt.Print(C.X)
	return resp
}

func end() {
	fmt.Println("The End.")
	os.Exit(1)
}

func Run(in io.Reader) {
	var resp string

	fmt.Println(C.R + "Do you like waffles? " + C.Y)
	resp = prompt(in)

	//if !strings.HasPrefix(strings.ToLower(resp), "y") {

	if !YES.MatchString(resp) {
		end()
	}

	fmt.Println(C.R + "Do you like pancakes? " + C.Y)

	resp = prompt(in)
	if !YES.MatchString(resp) {
		end()
	}

}
