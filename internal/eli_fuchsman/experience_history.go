package elifuchsman

import (
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Experience struct {
	Company        string `json:"company"`
	Position       string `json:"position"`
	EmploymentType string `json:"employment_type"`
	Address        string `json:"address"`
	Website        string `json:"website"`
	Start          string `json:"start"`
	End            string `json:"end"`
	Description    string `json:"description"`
}

type ExperienceHistory struct {
	History []*Experience
}

func (c *EliFuchsmanClient) ReturnExperienceHistory(tableName string) (*ExperienceHistory, error) {
	fields := log.Fields{"full_name": "EliFuchsman", "tableName": tableName}

	expHistory := &ExperienceHistory{History: make([]*Experience, 0)}
	dynamoExpHistory, err := c.edb.ReturnExperienceHistory(tableName)
	if err != nil {
		log.WithFields(fields).Errorf("ERROR FETCHING EXPERIENCE HISTORY FROM DYNAMODB: %+v", err)
		return nil, err
	}

	for _, exp := range dynamoExpHistory.History {
		newExp := &Experience{
			Company:        exp.Company,
			Position:       exp.Position,
			EmploymentType: exp.EmploymentType,
			Address:        exp.Address,
			Website:        exp.Website,
			Start:          exp.Start,
			End:            exp.End,
			Description:    exp.Description,
		}
		expHistory.History = append(expHistory.History, newExp)
	}

	sort.Slice(expHistory.History, func(i, j int) bool {
		return compareEndDates(expHistory.History[i].End, expHistory.History[j].End)
	})

	return expHistory, nil
}

// Helper to split the month and the year of the End attribute
func parseMonthYear(date string) (month, year string) {
	parts := strings.Fields(date)
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	return "", ""
}

// Helper to compare end dates
func compareEndDates(date1, date2 string) bool {
	month1, year1 := parseMonthYear(date1)
	month2, year2 := parseMonthYear(date2)

	if year1 != year2 {
		return year1 > year2
	}

	return strings.Index("JanuaryFebruaryMarchAprilMayJuneJulyAugustSeptemberOctoberNovemberDecember", month1) >
		strings.Index("JanuaryFebruaryMarchAprilMayJuneJulyAugustSeptemberOctoberNovemberDecember", month2)
}
