package comic

import (
    "bufio"
    "os"
    "strconv"
)

var bsize = 1
var idfile = "xkcd.id"
var contentfile = "xkcd.content"

func Refresh() {
    // read last ID saved to contentfile from idfile

    // incrementally download next comic record until not found (404)

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

func openFile(name string) (*os.File, error) {
    file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0755)
    if err != nil {
        return nil, err
    }
    return file, nil
}
