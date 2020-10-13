package console

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

type ConsoleReader interface {
	ReadLine(prompt string) (string, error)
	ReadPassword(prompt string) (string, error)
	ReadInt(prompt string) (int, error)
	Print(prompt string) error
	Println(prompt string) error
	EnableTty()
}

type DefaultConsoleReader struct {
	Tty bool
}

func NewConsoleReader() *DefaultConsoleReader {
	console := &DefaultConsoleReader{
		Tty: true,
	}
	return console
}

func (r DefaultConsoleReader) EnableTty() {
	r.Tty = true
	fmt.Fprint(os.Stdin, r.Tty)
}

func (r *DefaultConsoleReader) Print(prompt string) error {
	if r.Tty {
		// fmt.Fprint(os.Stdin, "Print tty")
		fmt.Fprint(os.Stdin, prompt)
	} else {
		// fmt.Fprint(os.Stdin, "Print")
		fmt.Fprint(os.Stderr, prompt)
	}
	return nil
}
func (r *DefaultConsoleReader) Println(prompt string) error {
	if r.Tty {
		// fmt.Fprintln(os.Stdin, "Fprintln tty")
		fmt.Fprintln(os.Stdin, prompt)
	} else {
		// fmt.Fprintln(os.Stdin, "Fprintln")
		fmt.Fprintln(os.Stderr, prompt)
	}
	return nil
}

func (r *DefaultConsoleReader) ReadLine(prompt string) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	r.Print(prompt)

	var s string
	scanner.Scan()
	if scanner.Err() != nil {
		return "", scanner.Err()
	}
	s = scanner.Text()
	return s, nil
}

func (r *DefaultConsoleReader) ReadInt(prompt string) (int, error) {
	var s string
	var err error
	if s, err = r.ReadLine(prompt); err != nil {
		return -1, err
	}

	var i int
	if i, err = strconv.Atoi(s); err != nil {
		return -1, err
	}

	return i, nil
}

func (r *DefaultConsoleReader) ReadPassword(prompt string) (string, error) {
	r.Print(prompt)
	pass, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	return string(pass[:]), nil
}
