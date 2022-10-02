package eightball

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/rwxrob/boost/weeks/20/cli"
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

// Respond will return a random response from list of Responses.
func Respond() string {
	rand.Seed(time.Now().UnixNano())
	return Responses[rand.Intn(len(Responses))]
}

// Run starts an interactive eightball session prompting the user for
// input and answering and then repeating until interrupted.
func Run() {
	fmt.Println(art)
	fmt.Println("Welcome to the magic eightball!")
	fmt.Println("(Enter your yes or no question)\n")
	for {
		cli.Prompt(os.Stdin, "🎱 ")
		fmt.Println(Respond())
	}
}
