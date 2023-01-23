package github

import "testing"

func TestCreate(t *testing.T) {
	createIssue(t)
}

func TestUpdate(t *testing.T) {
	issue := createIssue(t)

	issue.Body = "update"
	issue.Labels = append(issue.Labels, &Label{Name: "update"})
	var labels []string
	for _, l := range issue.Labels {
		labels = append(labels, l.Name)
	}

	issue, err := UpdateIssue(issue, "JasperLab", "go-programming-language-solutions")
	if err != nil {
		t.Fatal(err)
	}

	if issue.Body != "update" {
		t.Error("updated issue body != 'update'")
	}
	for i, l := range labels {
		if issue.Labels[i] == nil || issue.Labels[i].Name != l {
			t.Errorf("label %d != '%s'", i, l)
		}
	}

	issue.State = "closed"
	issue, err = UpdateIssue(issue, "JasperLab", "go-programming-language-solutions")
	if err != nil {
		t.Error("Failed to close the test issue due to " + err.Error())
	}
}

func TestClose(t *testing.T) {
}

func createIssue(t *testing.T) *Issue {
	i := &Issue{
		Title: "test",
		Body:  "body",
	}

	labels := [...]string{"label1", "label2"}
	for _, l := range labels {
		i.Labels = append(i.Labels, &Label{Name: l})
	}

	i.User = &User{
		Login: "JasperLab",
	}

	issue, err := CreateIssue("JasperLab", "go-programming-language-solutions", i)
	if err != nil {
		t.Fatal(err)
	}
	if issue.Title != "test" {
		t.Error("test title != 'Test'")
	}
	if issue.Body != "body" {
		t.Error("test body != 'Test'")
	}
	for i, l := range labels {
		if issue.Labels[i] == nil || issue.Labels[i].Name != l {
			t.Errorf("test label %d != %s", i, l)
		}
	}

	return issue
}

func deleteIssue(t *testing.T) {
}
