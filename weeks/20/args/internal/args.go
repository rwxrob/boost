package internal

import (
	"fmt"
	"strings"
)

func Output(args ...string) string {
	var buf string
	buf += strings.Join(args, " ") + "\n"
	for n, val := range args {
		buf += fmt.Sprintf("$%v --> %q\n", n, val)
	}
	return buf
}
