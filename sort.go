package notionapi

type SortOrder string

type TimestampType string

type SortObject struct {
	Property  string        `json:"property"`
	Timestamp TimestampType `json:"timestamp"`
	Direction SortOrder     `json:"direction"`
}
