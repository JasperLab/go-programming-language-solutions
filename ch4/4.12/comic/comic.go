package comic

import (
    "bufio"
	"encoding/json"
	"fmt"
	"net/http"
    "os"
    "strconv"
)

const size = 1
const idfile = "xkcd.id"
const ontentfile = "xkcd.content"
const xkcdUrl = "https://xkcd.com/%d/info.0.json"


func Refresh() (uint, error) {
    // read last ID saved to contentfile from idfile
    _, err := lastId()
    if err != nil {
        return 0, err
    }

    // incrementally download next comic record until not found (404)
	return 0, nil
}

func lastId() (uint, error) {
    f, err := openFile(idfile)
    if err != nil {
        return 0, err
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    text := "1"
    if scanner.Scan() {
        text = scanner.Text()
    } else if scanner.Err() != nil {
        // scanning ID error: we do not want tolerate any issues with the ID file
        return 0, err
    }

    id, err := strconv.Atoi(text)
    if err != nil || id <= 0 {
        // we can tolerate invalid index
        return 1, nil
    }
    return uint(id), nil
}

func updateId(id uint) error {
    str := strconv.FormatUint(uint64(id), 10)
    return os.WriteFile(idfile, []byte(str), 0)
}

func getRecord(id uint) (*Record, error) {
	url := fmt.Sprintf(xkcdUrl, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Record query failed: %s", resp.Status)
	}

	var result Record
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func openFile(name string) (*os.File, error) {
    file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0755)
    if err != nil {
        return nil, err
    }
    return file, nil
}
