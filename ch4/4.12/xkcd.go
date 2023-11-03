package main

import (
	"gopl/ch4/4.12/comic"
	"log"
)

func init() {
	log.Println("Loading xkcd records...")
	len, err := comic.Refresh()
	if err != nil {
		log.Fatalf("Failed to load xkcd records: %v", err)
	}
	log.Printf("Loaded %d records", len)
}

func main() {
}

func printUsage() {
	log.Fatal("Usage: go xkcd <search_term>")
}
