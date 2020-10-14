package mocks

type ConsoleReaderMock struct{}

func (r ConsoleReaderMock) ReadLine(prompt string) (string, error) {
	return prompt, nil
}

func (r ConsoleReaderMock) ReadInt(prompt string) (int, error) {
	return 0, nil
}

func (r ConsoleReaderMock) ReadPassword(prompt string) (string, error) {
	return prompt, nil
}

func (r ConsoleReaderMock) Println(prompt string) error {
	return nil
}

func (r ConsoleReaderMock) Print(prompt string) error {
	return nil
}
