package domain

import "github.com/Code0716/graphql-study/graph/model"

type Member model.Member

type GetMemberParams model.GetMemberParams

func (m Member) TableName() string {
	return "member"
}

type Pager model.Pager

// StatusCode member status model
type StatusCode int

const (
	StatusCodeFriend StatusCode = iota
	StatusCodeFamily
	StatusCodeBusiness
	StatusCodeOther
)

var StatusMap = map[StatusCode]model.Status{
	StatusCodeFriend:   "FRIEND",
	StatusCodeFamily:   "FAMILY",
	StatusCodeBusiness: "BUSINESS",
	StatusCodeOther:    "OTHER",
}
