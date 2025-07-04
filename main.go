package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"

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
	pathChan := make(chan string)
	resChan := make(chan pkg.ResultCounts, 10)
	var wg sync.WaitGroup
	var wgCounter sync.WaitGroup
	var count pkg.ResultCounts

	wgCounter.Add(1)
	go counter(resChan, &count, &wgCounter)

	for range 10 {
		wg.Add(1)
		go worker(pathChan, resChan, &wg)
	}

	if isInputFromPipe() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			pathChan <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		}
	} else if len(os.Args) > 1 {
		for _, path := range os.Args[1:] {
			pathChan <- path
		}
	} else {
		fmt.Fprintln(os.Stderr, "Usage: either pipe file paths into stdin or provide file paths as arguments.")
		os.Exit(1)
	}

	close(pathChan)
	wg.Wait()
	close(resChan)
	wgCounter.Wait()
	fmt.Printf("Line count: %d, Byte count: %d, Word count: %d\n", count.LineCount, count.ByteCount, count.WordCount)
}

func worker(ch chan string, resChan chan pkg.ResultCounts, wg *sync.WaitGroup) {
	defer wg.Done()
	for path := range ch {
		err := pkg.Counter(path, resChan)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func counter(ch chan pkg.ResultCounts, count *pkg.ResultCounts, wg *sync.WaitGroup) {
	defer wg.Done()
	for res := range ch {
		count.LineCount += res.LineCount
		count.WordCount += res.WordCount
		count.ByteCount += res.ByteCount
	}
}
