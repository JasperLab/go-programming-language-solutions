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

	if len(records) != 2 {
		t.Errorf("%d records expected in memory, found %d", 2, len(records))
	}
}

func TestBackfillDatabase(t *testing.T) {
	//prefill records so only first two are missing (should work for the next few years :))
	const max_records = 100000
	for i := uint(1); i <= max_records; i++ {
		records[i] = new(Record)
	}
	records[uint(1)], records[uint(2)], records[uint(1000)] = nil, nil, nil

	n, last := backfillDatabase(testFile)

	if n != 3 {
		t.Errorf("%d loaded records expected, got %d", 3, n)
	}

	if last != 1000 {
		t.Errorf("%d last loaded expected, got %d", 1000, last)
	}

	if len(records) != max_records {
		t.Errorf("%d records expected, got %d", max_records, len(records))
	}
}
