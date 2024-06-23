package router

import (
	"log"
	"time"
)

func DBTimeToTime(DBTime string) string {
	timeFromSQL, err := time.Parse("2006-01-02T15:04:05Z", DBTime)
	if err != nil {
		log.Fatal(err)
	}
	return timeFromSQL.Format("2006-01-02 15:04:05")
}
