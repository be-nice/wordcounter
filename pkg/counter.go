package pkg

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func Counter(w, b, l *int, path string) error {
	file, err := os.Open(path)
	if err != nil {
		errors.New("Error opening file")
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		*l++
		*b += len(scanner.Bytes()) + 1
		*w += len(strings.Fields(scanner.Text()))
	}

	return nil
}
