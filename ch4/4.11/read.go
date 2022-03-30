package main

import (
	"fmt"
	"gopl/ch4/4.11/github"
	"log"
	"os"
)

func main() {
	if len(os.Args[1:]) != 3 {
		printUsage()
	}

	issue, err := github.GetIssue(os.Args[1], os.Args[2], os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("#%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
	fmt.Println("=====================")
	fmt.Println(issue.Body)

}

func printUsage() {
	fmt.Println("Usage:\n\tread <owner> <repo> <id>")
	os.Exit(1)
}
