package comic

type Record struct {
	Num        int    `json:"num"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
}
