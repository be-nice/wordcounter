package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/be-nice/wordcounter/pkg"
)

func isInputFromPipe() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice == 0
}

func main() {
	var wordCount int
	var byteCount int
	var lineCount int
	if isInputFromPipe() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			pkg.Counter(&wordCount, &byteCount, &lineCount, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		}
	} else if len(os.Args) > 1 {
		for _, path := range os.Args[1:] {
			pkg.Counter(&wordCount, &byteCount, &lineCount, path)
		}
	} else {
		fmt.Fprintln(os.Stderr, "Usage: either pipe file paths into stdin or provide file paths as arguments.")
		os.Exit(1)
	}

	fmt.Printf("Line count: %d, Byte count: %d, Word count: %d\n", lineCount, byteCount, wordCount)
}
