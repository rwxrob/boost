package eightball

import (
	"fmt"
	"math/rand"
	"time"
)

const art = `
        ____
    ,dP9CGG88@b,
  ,IP  _   Y888@@b,
 dIi  (_)   G8888@b
dCII  (_)   G8888@@b
GCCIi     ,GG8888@@@
GGCCCCCCCGGG88888@@@
GGGGCCCGGGG88888@@@@...
Y8GGGGGG8888888@@@@P.....
 Y88888888888@@@@@P......
  Y8888888@@@@@@@P'......
    @@@@@@@@@P'.......
        """"........
`

var Responses = [...]string{
	"Yes",
	"No",
	"Maybe",
	"Never",
}

// Prompt prints the optional prompt (> by default) and returns the
// string entered by the user.
func Prompt(args ...string) string {
	var val string
	p := "> "
	if len(args) > 0 {
		p = args[0]
	}
	fmt.Print(p)
	val = greet.ReadLine()
	return val
}

// Respond will return a random response from list of Responses.
func Respond() string {
	rand.Seed(time.Now().UnixNano())
	return Responses[rand.Intn(len(Responses))]
}
