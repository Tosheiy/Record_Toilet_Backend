package router

import (
	"hello/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func updateRecordHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
	}
	var new_t_record model.TOILET_RECORD
	if err := c.ShouldBindJSON(&new_t_record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := db.Exec("UPDATE toilet_records SET description = ?, length = ?, location = ?, feeling = ? WHERE id = ?",
		new_t_record.Description, new_t_record.Length, new_t_record.Location, new_t_record.Feeling, id)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, result)
}
