package router

import (
	"hello/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func postRecordHandler(c *gin.Context) {
	var new_t_record model.TOILET_RECORD
	if err := c.ShouldBindJSON(&new_t_record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//ここでDefaultの設定をしておかないと全て入力されていなかったとき困る
	current_Time := time.Now()
	new_t_record.Created_at = current_Time.Format("2006-01-02 15:04:05")

	_, err := db.Exec("INSERT INTO toilet_records (description, created_at, length, location, feeling) VALUES (?, ?, ?, ?, ?)",
		new_t_record.Description, new_t_record.Created_at, new_t_record.Length, new_t_record.Location, new_t_record.Feeling)

	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusCreated, new_t_record)
}
