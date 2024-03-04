package elifuchsman

import log "github.com/sirupsen/logrus"

type Project struct {
	ProdName      string   `json:"production_name"`
	Repository    string   `json:"repository"`
	Image         string   `json:"image"`
	Description   string   `json:"description"`
	LearningGoals string   `json:"learning_goals"`
	TechStack     []string `json:"tech_stack"`
}

type Projects struct {
	Projects []*Project
}

func (c *EliFuchsmanClient) ReturnProjects(tableName string) (*Projects, error) {
	fields := log.Fields{"full_name": "EliFuchsman", "tableName": tableName}

	projects := &Projects{Projects: make([]*Project, 0)}
	dynamoProjects, err := c.edb.ReturnProjects(tableName)
	if err != nil {
		log.WithFields(fields).Errorf("ERROR FETCHING PROJECTS FROM DYNAMODB: %+v", err)
		return nil, err
	}

	for _, proj := range dynamoProjects.Projects {
		newProj := &Project{
			ProdName:      proj.ProdName,
			Repository:    proj.Repository,
			Image:         proj.Image,
			Description:   proj.Description,
			LearningGoals: proj.LearningGoals,
			TechStack:     proj.TechStack,
		}
		projects.Projects = append(projects.Projects, newProj)
	}

	return projects, nil
}
