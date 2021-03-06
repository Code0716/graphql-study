package domain

import "github.com/Code0716/graphql-study/graph/model"

// Person model
type Person model.Person

// GetPersonParams model
type GetPersonParams model.GetPersonParams

// PersonPager parer
type PersonPager model.PersonPager

// TableName get Person table name
func (m Person) TableName() string {
	return "person"
}
