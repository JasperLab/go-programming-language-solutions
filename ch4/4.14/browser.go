package main

import (
	"gopl/ch4/4.11/github"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	issues, err := fetch()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	users := uniqueUsers(issues)
	milestones := uniqueMilestones(issues)

	if err := render(w, issues, milestones, users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func fetch() ([]*github.Issue, error) {
	return github.ListIssues()
}

func uniqueUsers(issues []*github.Issue) (users []*github.User) {
	u := make(map[int]*github.User)

	for _, issue := range issues {
		if issue.User != nil {
			u[issue.User.Id] = issue.User
		}
	}

	for _, user := range u {
		users = append(users, user)
	}

	return
}

func uniqueMilestones(issues []*github.Issue) (milestones []*github.Milestone) {
	m := make(map[int]*github.Milestone)

	for _, issue := range issues {
		if issue.Milestone != nil {
			m[issue.Milestone.Id] = issue.Milestone
		}
	}

	for _, milestone := range m {
		milestones = append(milestones, milestone)
	}

	return
}

func render(w http.ResponseWriter, issues []*github.Issue, milestones []*github.Milestone, users []*github.User) error {
	w.Write([]byte("<html><body>"))
	if err := renderIssues(w, issues); err != nil {
		return err
	}

	w.Write([]byte("<br>"))

	if err := renderMilestones(w, milestones); err != nil {
		return err
	}

	w.Write([]byte("<br>"))

	if err := renderUsers(w, users); err != nil {
		return err
	}
	w.Write([]byte("</body></html>"))

	return nil
}

func renderIssues(w http.ResponseWriter, issues []*github.Issue) error {
	const isssueTemplate = `
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
