package elifuchsman

import (
	"encoding/json"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

type Summary struct {
	Summary string `json:"summary"`
}

func (c *EliFuchsmanClient) ReturnSummary(filePath string) (*Summary, error) {
	fields := log.Fields{"file_path": filePath}

	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.WithFields(fields).Errorf("ERROR READING JSON FILE: %+v", err)
		return nil, err
	}

	var summary *Summary
	err = json.Unmarshal(jsonData, &summary)
	if err != nil {
		log.WithFields(fields).Errorf("ERROR UNMARSHALING JSON: %+v", err)
		return nil, err
	}

	return summary, nil
}
