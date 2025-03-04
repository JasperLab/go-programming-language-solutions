package main

import (
	"fmt"
	"gopl/ch4/4.11/github"
	"html/template"
	"log"
	"net/http"
	"os"
)

var owner, repo string

func main() {
	if len(os.Args[1:]) != 2 {
		printUsage()
	}

	owner, repo = os.Args[1], os.Args[2]

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func printUsage() {
	fmt.Println("Usage:\n\tbrowser <owner> <repo>")
	os.Exit(1)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<html><body>"))

	issues, err := github.GetIssues(owner, repo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	if err = renderIssues(w, issues); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	milestones, err := github.GetMilestones(owner, repo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	if err = renderMilestones(w, milestones); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	users, err := github.GetAssignees(owner, repo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	if err = renderUsers(w, users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Write([]byte("</body></html>"))
}

func renderIssues(w http.ResponseWriter, issues []*github.Issue) error {
	const isssueTemplate = `
		<br>
		<table>
			<caption>Issues</caption>
			{{range .}}
				<tr>
					<td>
						<a href='{{.HTMLURL}}'>{{.Title}}</a>
					</td>
				</tr>
			{{end}}
		</table>
	`
	t := template.Must(template.New("issue").Parse(isssueTemplate))
	err := t.Execute(w, issues)
	return err
}

func renderMilestones(w http.ResponseWriter, milestones []*github.Milestone) error {
	const milestoneTemplate = `
		<br>
		<table>
			<caption>Milestones</caption>
			{{range .}}
				<tr>
					<td>
						<a href='{{.HTMLURL}}'>{{.Title}}</a>
					</td>
				</tr>
			{{end}}
		</table>
	`
	t := template.Must(template.New("milestone").Parse(milestoneTemplate))
	err := t.Execute(w, milestones)
	return err
}

func renderUsers(w http.ResponseWriter, users []*github.User) error {
	const userTemplate = `
		<br>
		<table>
			<caption>Users</caption>
			{{range .}}
				<tr>
					<td>
						<a href='{{.HTMLURL}}'>{{.Login}}</a>
					</td>
				</tr>
			{{end}}
		</table>
	`
	t := template.Must(template.New("user").Parse(userTemplate))
	err := t.Execute(w, users)
	return err
}
