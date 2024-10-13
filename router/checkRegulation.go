package router

import (
	"errors"
	"fmt"
	"hello/model"
	"time"
	"unicode/utf8"
)

func checkRegulation(new_t_record model.TOILET_RECORD) error {

	// Created_at が既に設定されているかチェック
	if new_t_record.Created_at != "" {
		// 入力された日付のフォーマットが正しいかチェック
		_, err := time.Parse("2006-01-02 15:04", new_t_record.Created_at)
		if err != nil {
			return err
		}
	} else {
		var err error
		// Created_at が入力されていない場合は現在の時間を設定
		new_t_record.Created_at, err = CreateNowTime()
		if err != nil {
			fmt.Println("Error loading location:", err)
			return err
		}
	}

	if utf8.RuneCountInString(new_t_record.Location_at) > 20 || utf8.RuneCountInString(new_t_record.Description) > 50 {
		return errors.New("文字数が多すぎます")
	}

	return nil
}