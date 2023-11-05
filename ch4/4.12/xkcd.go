package main

import (
	"gopl/ch4/4.12/comic"
	"log"
	"os"
)

func init() {
	log.Println("Loading xkcd records...")
	length, err := comic.Refresh()
	if err != nil {
		log.Fatalf("Failed to load xkcd records: %v", err)
	}
	log.Printf("Loaded %d records", length)
}

func main() {
	if len(os.Args) != 2 {
		printUsage()
	}

	term := os.Args[1]

	results := comic.Search(term)

	if len(results) == 0 {
		log.Println("No comics found.")
		os.Exit(0)
	}

	for link, transcript := range results {
		log.Println("----------------")
		log.Printf("%s\t%s\n", link, transcript)
		log.Println("----------------")
	}
}

func printUsage() {
	log.Fatal("Usage: go xkcd <search_term>")
}
