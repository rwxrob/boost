package termcolors

import (
	"fmt"
	"math/rand"
	"time"
)

const Black = "\033[30m"
const Red = "\033[31m"
const Green = "\033[32m"
const Yellow = "\033[33m"
const Blue = "\033[34m"
const Magenta = "\033[35m"
const Cyan = "\033[36m"
const White = "\033[37m"
const Reset = "\033[0m"
const Clear = "\033[H\033[2J"
const CurOff = "\033[?25l"
const CurOn = "\033[?25h"

// Rand returns a random color ANSI escape between Black (30) and White
// (37).
func Rand() string {
	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(7)
	return fmt.Sprintf("\033[3%vm", v)
}
