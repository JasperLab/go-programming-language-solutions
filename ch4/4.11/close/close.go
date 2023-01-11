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

	owner, repo, issue_id := os.Args[1], os.Args[2], os.Args[3]
	issue, err := github.GetIssue(owner, repo, issue_id)
	if err != nil {
		log.Fatal(err)
	}

	issue.State = "closed"
	issue, err = github.UpdateIssue(issue, owner, repo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("#%-5d %9.9s %.55s %.55s\n", issue.Number, issue.User.Login, issue.State, issue.Title)
	fmt.Println("=====================")
	fmt.Println(issue.Body)

}

func printUsage() {
	fmt.Println("Usage:\n\tclose <owner> <repo> <id>")
	os.Exit(1)
}
