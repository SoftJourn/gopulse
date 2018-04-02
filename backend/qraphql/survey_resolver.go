package graphql

import (
	"gopulse/backend/models"

	"github.com/graph-gophers/graphql-go"
)

type surveyResolver struct {
	entity *models.Survey
}

func (r *surveyResolver) ID() graphql.ID {
	return graphql.ID(r.entity.ID)
}

func (r *surveyResolver) Title() string {
	return r.entity.Title
}

func (r *surveyResolver) Description() string {
	return r.entity.Description
}

func (r *surveyResolver) Recipients() string {
	return r.entity.Recipients
}

func (r *surveyResolver) Questions() string {
	return r.entity.Questions
}

func (r *surveyResolver) Schedule() string {
	return r.entity.Schedule
}

func ResolveCreateSurvey() (result *surveyResolver) {
	survey, err := models.CreateSurvey(title, description)
	if err != nil {
		return nil
	}
	result = &surveyResolver{entity: survey}
	return
}
