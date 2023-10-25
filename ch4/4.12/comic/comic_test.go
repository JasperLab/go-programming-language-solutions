package comic

import (
	"encoding/json"
	"os"
	"testing"
)

const baseUrl = "https://xkcd.com/%d/info.0.json"
const testFile = "/tmp/xkcd.test"

func init() {
	os.Truncate(testFile, 0)
}

func TestFetchRecord(t *testing.T) {
	record, err := fetchRecord(1, baseUrl)
	if err != nil {
		t.Error(err)
	}
	if record.Num != 1 {
		t.Errorf("Record ID: expected 1, got %d", record.Num)
	}
}

func TestAppendRecord(t *testing.T) {
	record, err := fetchRecord(1, baseUrl)
	if err != nil {
		t.Error(err)
	}

	err = appendRecord(record, testFile)
	if err != nil {
		t.Error(err)
	}
}

func TestLoadDatabase(t *testing.T) {
	os.Truncate(testFile, 0)

	r1, err := fetchRecord(1, baseUrl)
	if err != nil {
		t.Error(err)
	}
	err = appendRecord(r1, testFile)
	if err != nil {
		t.Error(err)
	}

	r2, err := fetchRecord(2, baseUrl)
	if err != nil {
		t.Error(err)
	}
	err = appendRecord(r2, testFile)
	if err != nil {
		t.Error(err)
	}

	n, err := loadDatabase(testFile)
	if err != nil {
		t.Error(err)
	}
	if n != 2 {
		t.Fatalf("Expected 2 loaded records, got %d", n)
	}

	e1, _ := json.Marshal(r1)
	a1, _ := json.Marshal(records[1])
	if string(e1) != string(a1) {
		t.Errorf("%s\n\nexpected; received\n\n%s", e1, a1)
	}

	e2, _ := json.Marshal(r2)
	a2, _ := json.Marshal(records[2])
	if string(e2) != string(a2) {
		t.Errorf("%s\n\nexpected; received\n\n%s", e2, a2)
	}
}
