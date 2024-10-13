package model

type TOILET_RECORD struct {
    ID          string    `json:"id" dynamodbav:"id"`
    Description string `json:"description" dynamodbav:"description"`
    Created_at  string `json:"created_at" dynamodbav:"created_at"`
    Length_time      int    `json:"length_time" dynamodbav:"length_time"`
    Location_at    string `json:"location_at" dynamodbav:"location_at"`
    Feeling     int    `json:"feeling" dynamodbav:"feeling"`
	Uid         string `dynamodbav:"uid"`
}

// User はユーザー情報を表す構造体です
type USER struct {
	UTID string    `json:"utid" dynamodbav:"utid"`
	UID string `dynamodbav:"uid"`
	APIKEY string `json:"apikey" dynamodbav:"apikey"`
}