package model

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue" // これを追加
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func GetUserRecord(utid string) ([]USER, error) {
    input := &dynamodb.QueryInput{
        TableName: aws.String("user_table"),
        KeyConditionExpression: aws.String("utid = :utid"),
        ExpressionAttributeValues: map[string]types.AttributeValue{
            ":utid": &types.AttributeValueMemberS{Value: utid},
        },
    }

    result, err := svc.Query(context.TODO(), input)
    if err != nil {
        return nil, err
    }

    var userRecords []USER
    for _, item := range result.Items {
        var userRecord USER
        err = attributevalue.UnmarshalMap(item, &userRecord)
        if err != nil {
            return nil, err
        }
        userRecords = append(userRecords, userRecord)
    }

    return userRecords, nil
}

func CheckUserRecord(uid string) (*USER, error) {
	// GetItemのためのデータを作成
	input := &dynamodb.QueryInput{
		TableName:              aws.String("user_table"),
		IndexName:              aws.String("UIDIndex"), // GSIのインデックス名
		KeyConditionExpression: aws.String("uid = :uid"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":uid": &types.AttributeValueMemberS{Value: uid},
		},
	}

	// GetItem実行
	result, err := svc.Query(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	// アイテムが存在しない場合は適切なエラーを返す
	if len(result.Items) == 0 {
		return nil, fmt.Errorf("no record found for uid: %s", uid)
	}

	// アイテムをUSER構造体にマッピング
	var userRecord USER
	err = attributevalue.UnmarshalMap(result.Items[0], &userRecord)
	if err != nil {
		return nil, err
	}

	return &userRecord, nil
}

func InsertUserRecord(record USER) error {
	// PutItemのためのデータを作成
	input := &dynamodb.PutItemInput{
		TableName: aws.String("user_table"),
		Item: map[string]types.AttributeValue{
			"uid":    &types.AttributeValueMemberS{Value: record.UID},
			"utid":   &types.AttributeValueMemberS{Value: record.UTID},
			"apikey": &types.AttributeValueMemberS{Value: record.APIKEY},
		},
		ConditionExpression: aws.String("attribute_not_exists(uid)"), // uidが存在しないことを条件にする
	}

	// PutItem実行
	_, err := svc.PutItem(context.TODO(), input)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUserRecord(uid string, utid string) error {
	// DeleteItemのためのデータを作成
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String("user_table"),
		Key: map[string]types.AttributeValue{
			"uid":  &types.AttributeValueMemberS{Value: uid},  // パーティションキー
			"utid": &types.AttributeValueMemberS{Value: utid}, // ソートキー
		},
	}

	// DeleteItem実行
	_, err := svc.DeleteItem(context.TODO(), input)
	if err != nil {
		return err
	}

	return nil
}

func tableExists(tableName string) (bool, error) {
	output, err := svc.DescribeTable(context.TODO(), &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		if err.Error() == "ResourceNotFoundException: Table not found" {
			return false, nil // テーブルが存在しない
		}
		return false, err // 他のエラー
	}
	return output.Table.TableStatus == types.TableStatusActive, nil // テーブルが存在する
}