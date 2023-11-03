package comic

import (
    "bufio"
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

const size = 1
const contentfile = "/tmp/xkcd.content"
const xkcdUrl = "https://xkcd.com/%d/info.0.json"

var records = make(map[uint]*Record)
var latestId uint

func Refresh() (int, error) {
    _, err := loadDatabase(contentfile)
    if err != nil {
        return 0, err
    }

    _, _ = backfillDatabase(contentfile)

    return len(records), nil
}

// loads database file into memory in its entirety
// returns the total number of successfully loaded records
func loadDatabase(filename string) (uint, error) {

    data, err := os.ReadFile(filename)
    if err != nil {
        if _, err = os.Create(filename); err != nil {
            return 0, err
        }
        return 0, nil
    }

    r := bytes.NewReader(data)
    s := bufio.NewScanner(r)
    s.Split(bufio.ScanLines)
    var currentLine string
    sum := 0
    for s.Scan() {
        currentLine = s.Text()
        var record Record
        if err = json.Unmarshal([]byte(currentLine), &record); err != nil {
            //ignore individual unparseable lines
            continue
        }

        records[uint(record.Num)] = &record
        sum += 1
    }

    return uint(sum), nil
}

func backfillDatabase(filename string) (uint, uint) {
    id := uint(1)
    fetched := uint(0)
    last := uint(0)

    for {
        if records[id] != nil {
            id += 1
            continue
        }

        r, err := fetchRecord(id, xkcdUrl)
        if err != nil || r == nil {
            if id == 404 {
                //oddly the 404 record is missing, probably on purpose
                id += 1
                continue
            }
            break
        }
        err = appendRecord(r, contentfile)
        if err != nil {
            fmt.Printf("Error saving record: %v\n", err)
            break
        }
        records[id] = r
        last = id
        id += 1
        fetched += 1
    }

    return fetched, last
}

func fetchRecord(id uint, baseUrl string) (*Record, error) {
    url := fmt.Sprintf(baseUrl, id)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // naive error handling
    // can be improved to handle throttling/exp back-off retries
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("Record query failed: %s", resp.Status)
    }

    var result Record
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    return &result, nil
}

func appendRecord(record *Record, filename string) error {
    file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    value, err := json.Marshal(record)
    if err != nil {
        return err
    }

    writer := bufio.NewWriter(file)
    if _, err = writer.Write(value); err != nil {
        return err
    }
    if _, err = writer.WriteString("\n"); err != nil {
        return err
    }

    return writer.Flush()
}

func findRecord(num uint, filename string) (*Record, error) {
    return nil, nil
}
