package router

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func deleteRecordHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
	}

	result, err := db.Exec("DELETE FROM toilet_records WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, result)
}
