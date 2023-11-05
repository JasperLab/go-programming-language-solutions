package comic

import (
	"fmt"
	"strings"
)

func Search(term string) map[string]string {
	result := make(map[string]string)
	t := strings.ToLower(term)

	for _, r := range records {
		if s := strings.ToLower(r.Transcript); strings.Contains(s, t) {
			k := fmt.Sprintf(comicUrl, r.Num)
			result[k] = r.Transcript
		}
	}

	return result
}
