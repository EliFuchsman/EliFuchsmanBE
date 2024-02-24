package elifuchsmandb

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Client interface {
	ReturnBasicInfo(tableName string) (*BasicInfo, error)
}

type DBClient interface {
	GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
	Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
}

type EliFuchsmanDB struct {
	DynamoDB DBClient
}

type BasicInfo struct {
	FullName   string `json:"full_name" dynamodbav:"FullName"`
	FirstName  string `json:"first_name" dynamodbav:"FirstName"`
	LastName   string `json:"last_name" dynamodbav:"LastName"`
	City       string `json:"city" dynamodbav:"City"`
	State      string `json:"state" dynamodbav:"State"`
	Profession string `json:"profession" dynamodbav:"Profession"`
}

func NewEliFuchsmanDB(region string, endpoint string) (*EliFuchsmanDB, error) {
	awsConfig := aws.Config{
		Region:   aws.String(region),
		Endpoint: aws.String(endpoint),
	}

	if endpoint != "" {
		awsConfig.Endpoint = aws.String(endpoint)
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, err
	}

	dynamoDBClient := dynamodb.New(sess, &awsConfig)

	return &EliFuchsmanDB{
		DynamoDB: dynamoDBClient,
	}, nil
}

func (edb *EliFuchsmanDB) ReturnBasicInfo(tableName string) (*BasicInfo, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"FullName": {
				S: aws.String("EliFuchsman"),
			},
		},
	}

	result, err := edb.DynamoDB.GetItem(input)
	if err != nil {
		return nil, err
	}

	item := &BasicInfo{}
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		log.Println("Error unmarshalling item:", err)
		return nil, err
	}

	return item, nil
}
