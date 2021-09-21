package domain

import "github.com/Code0716/graphql-study/graph/model"

// CreateClassification params
type CreateClassification model.CreateClassification

// Classification model
type Classification model.Classification

// ClassificationsResponse model
type ClassificationsResponse []*model.Classification

// TableName get Classification table name
func (m Classification) TableName() string {
	return "classification"
}
