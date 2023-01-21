package github

import "testing"

func TestCreate(t *testing.T) {
	createIssue(t)
}

func TestUpdate(t *testing.T) {
}

func TestClose(t *testing.T) {

}

func TestDelete(t *testing.T) {
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

	return i
}

func deleteIssue(t *testing.T) {
}
