package console

import (
	"bufio"
	"os"
)

var input *bufio.Scanner

func init() {
	input = bufio.NewScanner(os.Stdin)
}

// ReadLine - Read input until an 'enter' key.
func ReadLine() string {
	input.Scan()
	return input.Text()
}

// Input - Writes the message and awaits for user input on the same line.
// Works just like the python function 'input()'.
// Combines the work of Write & Read into one call.
func Input(msg string) string {
	os.Stdout.WriteString(msg)
	return ReadLine()
}
