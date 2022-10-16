package termcolors

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"golang.org/x/term"
)

var (
	colorsOff bool
	Black     = "\033[30m"
	K         = Black
	Red       = "\033[31m"
	R         = Red
	Green     = "\033[32m"
	G         = Green
	Yellow    = "\033[33m"
	Y         = Yellow
	Blue      = "\033[34m"
	B         = Blue
	Magenta   = "\033[35m"
	M         = Magenta
	Cyan      = "\033[36m"
	C         = Cyan
	White     = "\033[37m"
	W         = White
	Reset     = "\033[0m"
	X         = Reset
	Clear     = "\033[H\033[2J"
	CurOff    = "\033[?25l"
	CurOn     = "\033[?25h"
)

// OnIfTerminal will activate the package colors if an interactive
// terminal is detected. Othewise, it disables them.
func OnIfTerminal(a any) {
	if fd, ok := a.(*os.File); !(ok && term.IsTerminal(int(fd.Fd()))) {
		Off()
		colorsOff = true
	} else {
		On()
		colorsOff = false
	}
}

func On() {
	Black = "\033[30m"
	K = Black
	Red = "\033[31m"
	R = Red
	Green = "\033[32m"
	G = Green
	Yellow = "\033[33m"
	Y = Yellow
	Blue = "\033[34m"
	B = Blue
	Magenta = "\033[35m"
	M = Magenta
	Cyan = "\033[36m"
	C = Cyan
	White = "\033[37m"
	W = White
	Reset = "\033[0m"
	X = Reset
	Clear = "\033[H\033[2J"
	CurOff = "\033[?25l"
	CurOn = "\033[?25h"
}

// Off sets all the package color variables to empty strings.
func Off() {
	Black = ""
	K = ""
	Red = ""
	R = ""
	Green = ""
	G = ""
	Yellow = ""
	Y = ""
	Blue = ""
	B = ""
	Magenta = ""
	M = ""
	Cyan = ""
	C = ""
	White = ""
	W = ""
	Reset = ""
	X = ""
	Clear = ""
	CurOff = ""
	CurOn = ""
}

// Rand returns a random color ANSI escape between Black (30) and White
// (37).
func Rand() string {
	if colorsOff {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(7)
	return fmt.Sprintf("\033[3%vm", v)
}
