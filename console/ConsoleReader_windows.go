// +build windows

package console

import (
	"bufio"
	"os"
)

func NewScanner() *bufio.Scanner {
	return bufio.NewScanner(os.Stdin)
}
