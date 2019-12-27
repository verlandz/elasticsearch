package models

type (
	ESResult struct {
		Took     int          `json:"took"`
		TimedOut bool         `json:"timed_out"`
		Shards   shardsResult `json:"_shards"`
		Hits     hitsResult   `json:'hits'`
	}
	shardsResult struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	}
	hitsResult struct {
		Total    totalResult  `json:"total"`
		MaxScore float64      `json:"max_score"`
		Hits     []hitsDetail `json:"hits"`
	}
	totalResult struct {
		Value    int    `json:"value"`
		Relation string `json:"relation"`
	}
	hitsDetail struct {
		Index  string      `json:"_index"`
		Type   string      `json:"_type"`
		ID     string      `json:"_id"`
		Score  float64     `json:"_score"`
		Source interface{} `json:"_source"`
	}
)
