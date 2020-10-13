package console

import (
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
	Println(prompt string) error
	SetDevice(device *os.File)
}

type DefaultConsoleReader struct {
	Device *os.File
}

func NewConsoleReader() DefaultConsoleReader {
	console := DefaultConsoleReader{
		Device: os.Stderr,
	}
	return console
}

func (r DefaultConsoleReader) SetDevice(device *os.File) {
	r.Device = device
}

func (r DefaultConsoleReader) Println(prompt string) error {
	fmt.Fprintln(r.Device, prompt)
	return nil
}

func (r DefaultConsoleReader) ReadLine(prompt string) (string, error) {
	scanner := NewScanner()
	fmt.Fprint(r.Device, prompt)
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
	fmt.Fprint(r.Device, prompt)
	pass, err := terminal.ReadPassword(int(syscall.Stdin))

	if err != nil {
		return "", err
	}

	return string(pass[:]), nil
}
