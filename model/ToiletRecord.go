package model

type TOILET_RECORD struct {
	ID          int    `json:"id`
	Description string `json:"description"`
	Created_at  string `json:"timestamp"`
	Length      int    `json:"length`
	Location    string `json:"location"`
	Feeling     int    `json:"feeling"`
	Uid         string
}

// User はユーザー情報を表す構造体です
type USER struct {
	ID     int
	UTID string    `json:"utid"`
	UID string
	APIKEY string `json:"apikey"`
}