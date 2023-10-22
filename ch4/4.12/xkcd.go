package main

import (
	"gopl/ch4/4.12/comic"
	"log"
)

func init() {
	comic.Refresh()
}

func main() {
}

func printUsage() {
	log.Fatal("Usage: go xkcd <search_term>")
}
