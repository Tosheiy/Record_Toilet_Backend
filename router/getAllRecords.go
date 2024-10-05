package router

import (
	"fmt"
	"hello/firebase_setting"
	"hello/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"strings"
)

func getAllRecordsHandler(c *gin.Context) {

	// Authorizationヘッダーの取得
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		return
	}

	// Bearer プレフィックスを取り除く
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		return
	}
	idToken := parts[1]

	// トークンの検証
	VerifiedToken, err := firebase_setting.VerifyIDToken(idToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Verified user id: %+v\n", VerifiedToken.UID)



	rows := model.QueryDB("SELECT * FROM toilet_records WHERE uid = ?", VerifiedToken.UID)
	defer rows.Close()

	var toilet_records []model.TOILET_RECORD
	for rows.Next() {
		var t_record model.TOILET_RECORD

		err := rows.Scan(&t_record.ID, &t_record.Description, &t_record.Created_at, &t_record.Length, &t_record.Location, &t_record.Feeling, &t_record.Uid)
		if err != nil {
			log.Fatal(err)
		}
		// とってきたstring型のTIMEをTime.time型に変換し、それをレイアウトを少し変えて文字列型に再変換している（Str(RFC3339形式) -> Time.time -> Str）
		t_record.Created_at = DBTimeToTime(t_record.Created_at)

		toilet_records = append(toilet_records, t_record)
	}
	c.JSON(http.StatusOK, toilet_records)
}
