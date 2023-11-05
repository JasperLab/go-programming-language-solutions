package comic

import (
	"os"
	"testing"
)

func TestSearch(t *testing.T) {
	os.Truncate(testFile, 0)
	clear(records)

	r2, err := fetchRecord(2, baseUrl)
	if err == nil {
		appendRecord(r2, testFile)
	}
	r1000, err := fetchRecord(1000, baseUrl)
	if err == nil {
		appendRecord(r1000, testFile)
	}
	if n, err := loadDatabase(testFile); err != nil {
		t.Fatalf("Error loading test database: %v", err)
	} else if n != 2 {
		t.Fatalf("Exepcted to load 2 records; loaded %d instead", n)
	}

	term := "  "
	results := Search(term)
	if len(results) != 0 {
		t.Fatalf("Expected %d, got %d for search term '%s'", 0, len(results), term)
	}

	term = "Petit"
	results = Search(term)
	if len(results) != 1 {
		t.Fatalf("Expected %d, got %d for search term '%s'", 1, len(results), term)
	}

	term = "WHICH"
	results = Search(term)
	if len(results) != 2 {
		t.Fatalf("Expected %d, got %d for search term '%s'", 2, len(results), term)
	}
}
