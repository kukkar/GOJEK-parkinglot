package lib

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Common Function with dependency of input injected via argument
func ReadAndProcessFromInput(input io.Reader) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		command := strings.ToLower(scanner.Text())
		processCommand(command)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// Function to read from stdin and process the command with arguments
func ReadAndProcessStdIn() {
	ReadAndProcessFromInput(os.Stdin)
}

// Function to read from file line by line and process the command with arguments
func ReadAndProcessFromFile(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	ReadAndProcessFromInput(f)
}
