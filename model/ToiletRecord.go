package model

type TOILET_RECORD struct {
	ID          int    `json:"id`
	Description string `json:"description"`
	Created_at  string `json:"timestamp"`
	Length      int    `json:"length`
	Location    string `json:"location"`
	Feeling     int    `json:"feeling"`
}