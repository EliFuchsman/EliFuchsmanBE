package elifuchsman

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

type BasicInfo struct {
	FullName   string `json:"full_name"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	City       string `json:"city"`
	State      string `json:"state"`
	Profession string `json:"profession"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}

func (c *EliFuchsmanClient) ReturnBasicInfo(filePath string) (*BasicInfo, error) {
	fields := log.Fields{"file_path": filePath}

	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.WithFields(fields).Errorf("ERROR READING JSON FILE: %+v", err)
		return nil, err
	}

	var basicInfo *BasicInfo
	err = json.Unmarshal(jsonData, &basicInfo)
	if err != nil {
		log.WithFields(fields).Errorf("ERROR UNMARSHALING JSON: %+v", err)
		return nil, err
	}

	return basicInfo, nil
}
