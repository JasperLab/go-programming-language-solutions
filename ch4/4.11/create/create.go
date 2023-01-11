package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) < 1 {
		printUsage()
	}

	title := os.Args[1]

	var assignees string[]
	if len(os.Args[1:]) > 1 {
		assignees = strings.Split(args[2])
	}

	var labels  string[]
	if len(os.Args[1:] > 2) {
		labels = strings.Split(args[3])
	}

		
}

func printUsage() {
	fmt.Println("Usage:\n\t go create/create.go --title='title' [--assignee=name1[,name2,...] --labels[=label1[,label2,...]]]")
	os.Exit(1)
}
