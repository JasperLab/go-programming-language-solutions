package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filenames)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			for filename, count := range filenames[line] {
				fmt.Printf("\t%s - %d\n", filename, count)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, filenames map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if filenames[input.Text()] == nil {
			filenames[input.Text()] = make(map[string]int)
		}
		filenames[input.Text()][f.Name()]++
	}
}
