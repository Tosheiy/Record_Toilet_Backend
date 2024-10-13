package model

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

var svc *dynamodb.Client

func ConnectDB() (*sql.DB, error) {
	// AWS SDKの設定を読み込む
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-northeast-1"))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// DynamoDBクライアントを初期化
	svc = dynamodb.NewFromConfig(cfg)

	// テスト
	tableName := "user_table"
	input := &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	}
	// DescribeTableでテーブル情報を取得
	result, err := svc.DescribeTable(context.TODO(), input)
	if err != nil {
		log.Fatal("使用するテーブルが作成されていますん")
		return nil, fmt.Errorf("failed to describe table: %v", err)
	}

	fmt.Printf("Table %s exists. Status: %s\n", tableName, string(result.Table.TableStatus))

	// InitDB()

	return nil, nil
}

func InitDB() {
	// NoSQL

	tableName := "toilet_records"

	// テーブルの存在確認
	exists, err := tableExists(tableName)
	if err != nil {
		log.Fatalf("failed to check if table exists, %v", err)
	}

	if exists {
		fmt.Printf("Table %s already exists.\n", tableName)
		return // すでにテーブルが存在する場合は何もしない
	}
	// テーブル作成リクエスト
	input := &dynamodb.CreateTableInput{
		TableName: aws.String("toilet_records"),
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("uid"),
				KeyType:       types.KeyTypeHash, // パーティションキー
			},
			{
				AttributeName: aws.String("id"),
				KeyType:       types.KeyTypeRange, // ソートキー
			},
		},
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("uid"),
				AttributeType: types.ScalarAttributeTypeS, // 文字列型
			},
			{
				AttributeName: aws.String("id"),
				AttributeType: types.ScalarAttributeTypeS, // 文字列型
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	// テーブルを作成
	_, err = svc.CreateTable(context.TODO(), input)
	if err != nil {
		log.Fatalf("failed to create table, %v", err)
	}

	fmt.Println("Table created successfully!")

	tableName = "user_table"

	// テーブルの存在確認
	exists, err = tableExists(tableName)
	if err != nil {
		log.Fatalf("failed to check if table exists, %v", err)
	}

	if exists {
		fmt.Printf("Table %s already exists.\n", tableName)
		return // すでにテーブルが存在する場合は何もしない
	}

	// テーブル作成リクエスト
	input = &dynamodb.CreateTableInput{
		TableName: aws.String("user_table"),
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("utid"), // パーティションキー
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("uid"), // ソートキー
				KeyType:       types.KeyTypeRange,
			},
		},
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("utid"),         // uidを属性定義に追加
				AttributeType: types.ScalarAttributeTypeS, // 文字列型
			},
			{
				AttributeName: aws.String("uid"),          // utidを属性定義に追加
				AttributeType: types.ScalarAttributeTypeS, // 文字列型
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		// Global Secondary Indexの設定
		GlobalSecondaryIndexes: []types.GlobalSecondaryIndex{
			{
				IndexName: aws.String("UIDIndex"), // インデックス名
				KeySchema: []types.KeySchemaElement{
					{
						AttributeName: aws.String("uid"), // uid を GSI のパーティションキーとして指定
						KeyType:       types.KeyTypeHash,
					},
				},
				Projection: &types.Projection{
					ProjectionType: types.ProjectionTypeAll, // 必要に応じて項目を選択
				},
				ProvisionedThroughput: &types.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(5),
					WriteCapacityUnits: aws.Int64(5),
				},
			},
		},
	}

	// テーブルを作成
	_, err = svc.CreateTable(context.TODO(), input)
	if err != nil {
		log.Fatalf("failed to create user_table, %v", err)
	}

	fmt.Println("user_table created successfully!")
}

func GenerateID() string {
	return uuid.NewString() // 新しいUUIDを生成
}
