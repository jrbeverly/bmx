package console

import (
	"bufio"
	"os"
)

func NewScanner() *bufio.Scanner {
	return bufio.NewScanner(os.Stdin)
}

func NewPromptWriter() os.File {
	return os.Stderr
}
