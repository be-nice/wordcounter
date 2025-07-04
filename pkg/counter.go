package pkg

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

type ResultCounts struct {
	WordCount int
	LineCount int
	ByteCount int
}

func Counter(path string, res chan ResultCounts) error {
	file, err := os.Open(path)
	if err != nil {
		errors.New("Error opening file")
		return err
	}
	defer file.Close()

	var counts ResultCounts
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		counts.WordCount += len(strings.Fields(scanner.Text()))
		counts.LineCount++
		counts.ByteCount += len(scanner.Bytes()) + 1

	}

	res <- counts
	return nil
}
