package elifuchsmandb

import "github.com/aws/aws-sdk-go/service/dynamodb"

type MockDBClient struct {
	GetItemOutput *dynamodb.GetItemOutput
	GetItemError  error

	QueryData *dynamodb.QueryOutput
	QueryErr  error
}

// Used to mock a DynamoDB request
func (m *MockDBClient) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return m.GetItemOutput, m.GetItemError
}

func (m *MockDBClient) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	return m.QueryData, m.QueryErr
}
