package elifuchsmandb

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	log "github.com/sirupsen/logrus"
)

type Client interface {
	ReturnEducationHistory(tableName string) (*EducationHistory, error)
	ReturnExperienceHistory(tableName string) (*ExperienceHistory, error)
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

type Education struct {
	FullName      string `json:"full_name" dynamodbav:"FullName"`
	EducationType string `json:"education_type" dynamodbav:"EducationType"`
	Name          string `json:"name" dynamodbav:"Name"`
	City          string `json:"city" dynamodbav:"City"`
	State         string `json:"state" dynamodbav:"State"`
	Degree        string `json:"degree" dynamodbav:"Degree"`
	Major         string `json:"major" dynamodbav:"Major"`
	From          string `json:"from" dynamodbav:"From"`
	To            string `json:"to" dynamodbav:"To"`
}

type EducationHistory struct {
	History []*Education
}

type Experience struct {
	FullName       string `json:"full_name" dynamodbav:"FullName"`
	ExperienceId   string `json:"experience_id" dynamodbav:"ExperienceId"`
	Company        string `json:"company" dynamodbav:"Company"`
	Position       string `json:"position" dynamodbav:"Position"`
	EmploymentType string `json:"employment_type" dynamodbav:"EmploymentType"`
	Address        string `json:"address" dynamodbav:"Address"`
	Website        string `json:"website" dynamodbav:"Website"`
	Start          string `json:"start" dynamodbav:"Start"`
	End            string `json:"end" dynamodbav:"End"`
	Description    string `json:"description" dynamodbav:"Description"`
}

type ExperienceHistory struct {
	History []*Experience
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

func (edb *EliFuchsmanDB) ReturnEducationHistory(tableName string) (*EducationHistory, error) {
	if tableName == "" {
		return nil, errors.New("tableName is required")
	}

	fullName := "EliFuchsman"
	input := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		KeyConditionExpression: aws.String("FullName = :fullName"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":fullName": {
				S: aws.String(fullName),
			},
		},
	}

	result, err := edb.DynamoDB.Query(input)
	if err != nil {
		log.WithError(err).Error("Error querying DynamoDB")
		return nil, fmt.Errorf("error querying DynamoDB: %w", err)
	}

	edHistory := &EducationHistory{History: make([]*Education, 0)}

	for _, item := range result.Items {
		ed := &Education{}
		if err = dynamodbattribute.UnmarshalMap(item, ed); err != nil {
			log.WithError(err).Error("Error unmarshaling item")
			log.WithField("rawItem", item).Error("Raw DynamoDB Item")
			return nil, fmt.Errorf("error unmarshaling item: %w", err)
		}
		edHistory.History = append(edHistory.History, ed)
	}
	return edHistory, nil
}

func (edb *EliFuchsmanDB) ReturnExperienceHistory(tableName string) (*ExperienceHistory, error) {
	if tableName == "" {
		return nil, errors.New("tableName is required")
	}

	fullName := "EliFuchsman"
	input := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		KeyConditionExpression: aws.String("FullName = :fullName"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":fullName": {
				S: aws.String(fullName),
			},
		},
	}

	result, err := edb.DynamoDB.Query(input)
	if err != nil {
		log.WithError(err).Error("Error querying DynamoDB")
		return nil, fmt.Errorf("error querying DynamoDB: %w", err)
	}

	expHistory := &ExperienceHistory{History: make([]*Experience, 0)}

	for _, item := range result.Items {
		ed := &Experience{}
		if err = dynamodbattribute.UnmarshalMap(item, ed); err != nil {
			log.WithError(err).Error("Error unmarshaling item")
			log.WithField("rawItem", item).Error("Raw DynamoDB Item")
			return nil, fmt.Errorf("error unmarshaling item: %w", err)
		}
		expHistory.History = append(expHistory.History, ed)
	}
	return expHistory, nil
}
