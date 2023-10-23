package comic

import "testing"

func TestLastId(t *testing.T) {
	lid, err := lastId()
	if err != nil {
		t.Errorf("Error reading last Id: %s", err.Error())
	}
	if lid <= 0 {
		t.Errorf("Invalid last Id: %d", lid)
	}
}

func TestUpdateId(t *testing.T) {
	currId, err := lastId()
	defer updateId(currId);
	nextId := currId + 1
	updateId(nextId)
	lastId, err := lastId()
	if err != nil {
		t.Errorf("Error reading last Id: %s", err.Error())
	}
	if lastId != nextId {
		t.Errorf("Expected: %d, read: %d", nextId, lastId)
	}
}

func TestGetRecord(t *testing.T) {
	record, err := getRecord(1)
	if err != nil {
		t.Error(err)
	}
	if record.Num != 1 {
		t.Errorf("Record ID: expected 1, got %d", record.Num)
	}
}
