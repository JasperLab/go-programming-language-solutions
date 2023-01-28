package main

import (
	"fmt"
	"gopl/ch4/4.11/github"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args[1:]) < 3 {
		printUsage()
	}
	args := os.Args[1:]
	owner := args[0]
	repo := args[1]
	issue_id := args[2]

	issue, err := github.GetIssue(owner, repo, issue_id)
	if err != nil {
		log.Fatal(err)
	}

	// handle assignees - 4th argument
	if len(args) > 3 {
		split := strings.Split(os.Args[1:][1], ",")
		for _, l := range split {
			assignee := &github.Assignee{Login: l}
			issue.Assignees = append(issue.Assignees, assignee)
		}
	}

	// handle labels - 5th argument
	if len(args) > 4 {
		split := strings.Split(os.Args[1:][2], ",")
		for _, l := range split {
			label := &github.Label{Name: l}
			issue.Labels = append(issue.Labels, label)
		}
	}

	//handle issue body via vi
	editor := "vi"
	if s := os.Getenv("EDITOR"); s != "" {
		editor = s
	}
	tmpDir := os.TempDir()
	tmpFile, err := ioutil.TempFile(tmpDir, "tempFilePrefix")
	if err != nil {
		log.Fatal(err)
	}
	_, err = tmpFile.WriteString(issue.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	cmd := exec.Command("sh", "-c", editor+" "+tmpFile.Name())
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
	b, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		log.Fatal(err)
	}
	issue.Body = string(b)

	issue, err = github.UpdateIssue(owner, repo, issue)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("#%-5d %9.9s %.55s %.55s\n", issue.Number, issue.User.Login, issue.State, issue.Title)
	fmt.Println("=====================")
	fmt.Println(issue.Body)
}

func printUsage() {
	fmt.Println("Usage:\n\t go update/update.go <owner> <repo> <issue_id> [<login1[,login2,...]>] [<label1[,label2,...]>]")
	os.Exit(1)
}
