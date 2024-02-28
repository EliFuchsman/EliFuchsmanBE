package elifuchsman

import log "github.com/sirupsen/logrus"

type Education struct {
	EducationType string `json:"education_type"`
	Name          string `json:"name"`
	City          string `json:"city"`
	State         string `json:"state"`
	Degree        string `json:"degree"`
	Major         string `json:"major"`
	From          string `json:"from"`
	To            string `json:"to"`
}

type EducationHistory struct {
	History []*Education
}

func (c *EliFuchsmanClient) ReturnEducationHistory(tableName string) (*EducationHistory, error) {
	fields := log.Fields{"full_name": "EliFuchsman", "tableName": tableName}

	edHistory := &EducationHistory{History: make([]*Education, 0)}
	dynamoEdHistory, err := c.edb.ReturnEducationHistory(tableName)
	if err != nil {
		log.WithFields(fields).Errorf("ERROR FETCHING EDUCATION HISTORY FROM DYNAMODB: %+v", err)
		return nil, err
	}

	for _, ed := range dynamoEdHistory.History {
		newEd := &Education{
			EducationType: ed.EducationType,
			Name:          ed.Name,
			City:          ed.City,
			State:         ed.State,
			Degree:        ed.Degree,
			Major:         ed.Major,
			From:          ed.From,
			To:            ed.To,
		}
		edHistory.History = append(edHistory.History, newEd)
	}

	return edHistory, nil
}
