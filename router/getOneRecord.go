package router

import (
	"hello/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getOneRecordHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
	}

	rows := model.QueryDB(db, "SELECT * FROM toilet_records WHERE id = ?", id)
	defer rows.Close()

	var toilet_records []model.TOILET_RECORD
	for rows.Next() {
		var t_record model.TOILET_RECORD

		err = rows.Scan(&t_record.ID, &t_record.Description, &t_record.Created_at, &t_record.Length, &t_record.Location, &t_record.Feeling)
		if err != nil {
			log.Fatal(err)
		}
		t_record.Created_at = DBTimeToTime(t_record.Created_at)
		toilet_records = append(toilet_records, t_record)
	}
	c.JSON(http.StatusOK, toilet_records)
}