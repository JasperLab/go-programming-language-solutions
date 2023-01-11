package github

import "testing"

func TestCreate(t *testing.T) {
	label1 := &Label{Name: "Test1"}
	label2 := &Label{Name: "Test2"}
	labels := [2]*Label{label1, label2}
	user := &User{
		Login: "JasperLab",
	}
	i := &Issue{
		Title:  "Test",
		Body:   "Test",
		Labels: labels[:],
		User:   user,
	}

	i, err := CreateIssue("JasperLab", "go-programming-language-solutions", i)
	if err != nil {
		t.Fatal(err)
	}
	if i.Title != "Test" {
		t.Error("test title != 'Test'")
	}
	if i.Body != "Test" {
		t.Error("test body != 'Test'")
	}
	if len(i.Labels) != 2 || i.Labels[0].Name != "Test1" || i.Labels[1].Name != "Test2" {
		t.Error("test labels != ['Test1', 'Test2']")
	}
}
