package elifuchsman

import (
	log "github.com/sirupsen/logrus"
)

type BasicInfo struct {
	FullName    string `json:"full_name"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	City        string `json:"city"`
	State       string `json:"state"`
	DateOfBirth string `json:"date_of_birth"`
}

func (c *EliFuchsmanClient) ReturnBasicInfo(tableName string) (*BasicInfo, error) {
	fields := log.Fields{"table_name": tableName}

	dynamoInfo, err := c.edb.ReturnBasicInfo(tableName)
	if err != nil {
		log.WithFields(fields).Errorf("ERROR FETCHING INFO FROM DYNAMODB: %+v", err)
		return nil, err
	}

	basicInfo := &BasicInfo{
		FullName:    dynamoInfo.FullName,
		FirstName:   dynamoInfo.FirstName,
		LastName:    dynamoInfo.LastName,
		City:        dynamoInfo.LastName,
		State:       dynamoInfo.State,
		DateOfBirth: dynamoInfo.LastName,
	}

	return basicInfo, nil
}
