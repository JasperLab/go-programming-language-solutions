package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("USAGE: go run count_words.go <file1> <file2>....<file_n>")
	}
	m := count_words(os.Args[1:])
	fmt.Println(sorted_counts_str(m))
}

func count_words(files []string) map[string]uint {
	counts := make(map[string]uint)
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			log.Fatalf("count_words: %v", err)
		}
		input := bufio.NewScanner(f)
		input.Split(bufio.ScanWords)
		for input.Scan() {
			counts[input.Text()]++
		}
	}
	return counts
}

func sorted_counts_str(m map[string]uint) string {
	var words []string
	var result []byte
	for word := range m {
		words = append(words, word)
	}
	sort.Strings(words)
	for _, word := range words {
		result = append(result, fmt.Sprintf("%v:\t%d", word, m[word])...)
	}
	return string(result)
}
