package cli

import (
	"bufio"
	"io"
	"strings"
)

// ReadLine takes any io.Reader and returns a trimmed string (initial
// and trailing white space) or an empty string and error if any error
// is encountered.
func ReadLine(in io.Reader) (string, error) {
	out, err := bufio.NewReader(in).ReadString('\n')
	out = strings.TrimSpace(out)
	return out, err
}
