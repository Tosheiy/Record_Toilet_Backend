package router

import (
	"fmt"
	"log"
	"time"
)

func DBTimeToTime(DBTime string) string {
	timeFromSQL, err := time.Parse("2006-01-02 15:04:00", DBTime)
	if err != nil {
		log.Fatal(err)
	}
	return timeFromSQL.Format("2006-01-02 15:04")
}


func CreateNowTime() (string, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return "", err
	}
	current_Time := time.Now().In(jst)
	return current_Time.Format("2006-01-02 15:04"), nil
}