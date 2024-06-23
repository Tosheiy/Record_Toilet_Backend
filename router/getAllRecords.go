package router

import (
	"hello/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllRecordsHandler(c *gin.Context) {
	rows := model.QueryDB(db, "SELECT * FROM toilet_records")
	defer rows.Close()

	var toilet_records []model.TOILET_RECORD
	for rows.Next() {
		var t_record model.TOILET_RECORD

		err := rows.Scan(&t_record.ID, &t_record.Description, &t_record.Created_at, &t_record.Length, &t_record.Location, &t_record.Feeling)
		if err != nil {
			log.Fatal(err)
		}
		// とってきたstring型のTIMEをTime.time型に変換し、それをレイアウトを少し変えて文字列型に再変換している（Str(RFC3339形式) -> Time.time -> Str）
		t_record.Created_at = DBTimeToTime(t_record.Created_at)

		toilet_records = append(toilet_records, t_record)
	}
	c.JSON(http.StatusOK, toilet_records)
}
