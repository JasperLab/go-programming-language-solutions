package main

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func TestEmptyHtml(t *testing.T) {
	content := ""

	reader := strings.NewReader(content)

	//NOTE: looks loke Parse(stream) returns <html><head></head><body></body></html>
	//		even for empty strings
	doc, err := html.Parse(reader)
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
	}

	nodes := make(map[string]int)
	result := count(nodes, doc)

	if len(result) != 3 {
		t.Errorf("Expected lenght: 3, got: %d\n\tMap:\n\t%v", len(result), result)
	}

	if result["html"] != 1 {
		t.Errorf("Expected 1 <html>, got %d", result["html"])
	}

	if result["head"] != 1 {
		t.Errorf("Expected 1 <head>, got %d", result["html"])
	}

	if result["body"] != 1 {
		t.Errorf("Expected 1 <html>, got %d", result["html"])
	}
}

func TestHtml(t *testing.T) {
	content := `
		<html>
			<head>Header</head>
			<body>
				<p>Paragraph with <b>bold</b> text</p>
				<p/>
			</body>
		</html>
	`

	reader := strings.NewReader(content)

	doc, err := html.Parse(reader)
	if err != nil {
		t.Errorf("Parsing error: %v\n", err)
	}
	nodes := make(map[string]int)
	result := count(nodes, doc)

	if len(result) != 5 {
		t.Errorf("Expected 5 distinct html elements, got %d\n\tMap:\n\t%v", len(result), result)
	}

	if result["html"] != 1 {
		t.Errorf("Expected 1 <html>, got %d", result["html"])
	}

	if result["head"] != 1 {
		t.Errorf("Expected 1 <head>, got %d", result["html"])
	}

	if result["body"] != 1 {
		t.Errorf("Expected 1 <html>, got %d", result["html"])
	}

	if result["p"] != 2 {
		t.Errorf("Expected 1 <html>, got %d", result["html"])
	}

	if result["b"] != 1 {
		t.Errorf("Expected 1 <html>, got %d", result["html"])
	}
}

func TestInvalidHtml(t *testing.T) {
	content := `
		<html>
			<head>Header</head>
			<body>
	`

	reader := strings.NewReader(content)

	doc, err := html.Parse(reader)
	if err != nil {
		t.Errorf("Unexpected error: %v\n", err)
	}

	nodes := make(map[string]int)
	result := count(nodes, doc)

	if len(result) != 3 {
		t.Errorf("Expected lenght: 3, got: %d\n\tMap:\n\t%v", len(result), result)
	}

	if result["html"] != 1 {
		t.Errorf("Expected 1 <html>, got %d", result["html"])
	}

	if result["head"] != 1 {
		t.Errorf("Expected 1 <head>, got %d", result["html"])
	}

	if result["body"] != 1 {
		t.Errorf("Expected 1 <html>, got %d", result["html"])
	}

}
