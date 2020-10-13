package console

import (
	"fmt"
	"strconv"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

type ConsoleReader interface {
	ReadLine(prompt string) (string, error)
	ReadPassword(prompt string) (string, error)
	ReadInt(prompt string) (int, error)
	Prompt(prompt string) error
	Println(prompt string) error
	Print(prompt string) error
}

type DefaultConsoleReader struct{}

func (r DefaultConsoleReader) Prompt(prompt string) error {
	outf := NewWriter()
	fmt.Fprintln(outf, prompt)
	return nil
}

func (r DefaultConsoleReader) Print(prompt string) error {
	outf := NewWriter()
	fmt.Fprint(outf, prompt)
	return nil
}

func (r DefaultConsoleReader) Println(prompt string) error {
	outf := NewWriter()
	fmt.Fprintln(outf, prompt)
	return nil
}

func (r DefaultConsoleReader) ReadLine(prompt string) (string, error) {
	scanner := NewScanner()
	outf := NewWriter()
	fmt.Fprint(outf, prompt)
	var s string
	scanner.Scan()
	if scanner.Err() != nil {
		return "", scanner.Err()
	}
	s = scanner.Text()
	return s, nil
}

func (r DefaultConsoleReader) ReadInt(prompt string) (int, error) {
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

func (r DefaultConsoleReader) ReadPassword(prompt string) (string, error) {
	outf := NewWriter()
	fmt.Fprint(outf, prompt)
	pass, err := terminal.ReadPassword(int(syscall.Stdin))

	if err != nil {
		return "", err
	}

	return string(pass[:]), nil
}
