package elifuchsman

import (
	edb "github.com/efuchsman/EliFuchsmanBE/internal/eli_fuchsman_db"
)

type Client interface {
	ReturnBasicInfo(filePath string) (*BasicInfo, error)
	ReturnEducationHistory(tableName string) (*EducationHistory, error)
	ReturnSummary(filePath string) (*Summary, error)
}

type EliFuchsmanClient struct {
	edb edb.Client
}

func NewEliFuchsmanClient(db edb.Client) *EliFuchsmanClient {
	return &EliFuchsmanClient{
		edb: db,
	}
}
