package model

import (
	"context"
	"fmt"
	"log"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue" // これを追加
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func GetAllRecordsDB(uid string) ([]TOILET_RECORD, error) {
	// Query inputの設定
	input := &dynamodb.QueryInput{
		TableName:              aws.String("toilet_records"),
		KeyConditionExpression: aws.String("uid = :uid"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":uid": &types.AttributeValueMemberS{Value: uid},
		},
	}

	// Query実行
	result, err := svc.Query(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	// アイテムをTOILET_RECORDにマッピング
	records := []TOILET_RECORD{}
	for _, item := range result.Items {
		var record TOILET_RECORD
		err = attributevalue.UnmarshalMap(item, &record)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func GetOneRecordDB(uid string, id string) (*TOILET_RECORD, error) {
	// Query inputの設定
	input := &dynamodb.QueryInput{
		TableName:              aws.String("toilet_records"),
		KeyConditionExpression: aws.String("uid = :uid AND id = :id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":uid": &types.AttributeValueMemberS{Value: uid},
			":id":  &types.AttributeValueMemberS{Value: id},
		},
	}

	// Query実行
	result, err := svc.Query(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	// レコードが見つからない場合は nil を返す
	if len(result.Items) == 0 {
		return nil, nil // or return an error if needed
	}

	// 最初のアイテムをTOILET_RECORDにマッピング
	var record TOILET_RECORD
	err = attributevalue.UnmarshalMap(result.Items[0], &record)
	if err != nil {
		return nil, err
	}

	return &record, nil // 構造体のポインタを返す
}

func InsertRecordsDB(new_record TOILET_RECORD) error {
	// Insertのためのデータを作成
	input := &dynamodb.PutItemInput{
		TableName: aws.String("toilet_records"),
		Item: map[string]types.AttributeValue{
			"uid":         &types.AttributeValueMemberS{Value: new_record.Uid},
			"id":          &types.AttributeValueMemberS{Value: new_record.ID}, // idを追加
			"created_at":  &types.AttributeValueMemberS{Value: new_record.Created_at},
			"description": &types.AttributeValueMemberS{Value: new_record.Description},
			"length_time": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", new_record.Length_time)}, //文字列に変換する必要あり
			"location_at":    &types.AttributeValueMemberS{Value: new_record.Location_at},
			"feeling":     &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", new_record.Feeling)}, //文字列に変換する必要あり
		},
	}

	// PutItem実行
	_, err := svc.PutItem(context.TODO(), input)
	if err != nil {
		log.Println("Error in PutItem:", err)
		return err
	}

	return nil
}

// レコードをアップデートする関数
func UpdateRecordDB(uid string, id string, updatedRecord TOILET_RECORD) error {
	// UpdateItemのためのリクエスト作成
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("toilet_records"),
		Key: map[string]types.AttributeValue{
			"uid": &types.AttributeValueMemberS{Value: uid},
			"id":  &types.AttributeValueMemberS{Value: id},
		},
		UpdateExpression: aws.String("SET created_at = :created_at, description = :description, length_time = :length_time, location_at = :location_at, feeling = :feeling"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":created_at":  &types.AttributeValueMemberS{Value: updatedRecord.Created_at},
			":description": &types.AttributeValueMemberS{Value: updatedRecord.Description},
			":length_time": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", updatedRecord.Length_time)},
			":location_at":    &types.AttributeValueMemberS{Value: updatedRecord.Location_at},
			":feeling":     &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", updatedRecord.Feeling)},
		},
	}

	// UpdateItem実行
	_, err := svc.UpdateItem(context.TODO(), input)
	if err != nil {
		return err
	}

	return nil
}

func DeleteRecordsDB(uid string, id string) error {
	// DeleteItemリクエスト
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String("toilet_records"),
		Key: map[string]types.AttributeValue{
			"uid": &types.AttributeValueMemberS{Value: uid}, // パーティションキー
			"id":  &types.AttributeValueMemberS{Value: id},  // ソートキー
		},
	}

	// アイテムを削除
	_, err := svc.DeleteItem(context.TODO(), input)
	if err != nil {
		log.Println("Error in DeleteItem:", err)
		return err
	}

	return nil
}