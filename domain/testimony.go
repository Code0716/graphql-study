package domain

import (
	"github.com/Code0716/graphql-study/graph/model"
)

// Testimony is Testimony model
type Testimony model.Testimony

// TableName get Testimony table name
func (m Testimony) TableName() string {
	return "testimony"
}

// GetTestimonyParams get Testimony params
type GetTestimonyParams model.GetTestimonyParams

type TestimonyStatus model.TestimonyStatus

const (
	TestimonyStatusDraft  = "DRAFT"
	TestimonyStatusPublic = "PUBLIC"
	TestimonyStatusOther  = "OTHER"
)

func GetGetTestimonyStatus(status string) string {
	switch status {
	case TestimonyStatusDraft:
		return TestimonyStatusDraft
	case TestimonyStatusPublic:
		return TestimonyStatusPublic
	case TestimonyStatusOther:
		return TestimonyStatusOther
	default:
		return TestimonyStatusDraft
	}
}
