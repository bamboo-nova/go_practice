package main

// Flowchart1
// make main.go and main_test.go
// `go test -v` and `go build`->wc executable in the directory
// `echo "My first command line tool with Go" | ./wc`

// Flowchart2 (Adding Command-Line Flags)
// add package `flag`
// cat main.go | ./wc -l

// Flowchart3 (compilling tools for different platforms)
// GOOS=windows go build

// Exercise
// count bytes (cat main.go | ./wc -b)

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// Defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	// Defining a boolean flag -l to count lines instead of words
	bytes := flag.Bool("b", false, "Count lines")
	// Parsing the flags provided by the user
	flag.Parse()

	// Calling the count function to count the number of words
	// received from the Standard Input and printing it out
	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// If the count lines flag is not set, we want to count words so we define
	// the scanner split type to words (default is split by lines)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	// Split by bytes
	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	// Defining a counter
	wc := 0

	// For every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	// Return the total
	return wc
}
