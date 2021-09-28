// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Classification struct {
	ID        string     `json:"id"`
	ClassID   string     `json:"classId"`
	ClassName string     `json:"class_name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type CommonResponse struct {
	Message string `json:"Message"`
}

type CreateClassification struct {
	ClassID   *string `json:"class_id"`
	ClassName string  `json:"class_name"`
}

type CreatePerson struct {
	ID          *int    `json:"id"`
	PersonID    *string `json:"person_id"`
	Name        string  `json:"name"`
	Address     *string `json:"address"`
	PhoneNumber *string `json:"phone_number"`
	ClassName   *string `json:"class_name"`
}

type GetPersonParams struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}

type GetTestimonyParams struct {
	TestimonyID *string    `json:"testimony_id"`
	PersonID    *string    `json:"person_id"`
	Status      *string    `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
}

type Pager struct {
	Limit  *int `json:"limit"`
	Offset *int `json:"offset"`
}

type Person struct {
	ID          int        `json:"id"`
	PersonID    string     `json:"person_id"`
	Name        string     `json:"name"`
	Address     string     `json:"address"`
	PhoneNumber string     `json:"phone_number"`
	ClassName   string     `json:"class_name"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type PersonPager struct {
	Limit     *int    `json:"limit"`
	Offset    *int    `json:"offset"`
	ClassName *string `json:"class_name"`
}

type Testimony struct {
	ID          int        `json:"id"`
	TestimonyID string     `json:"testimony_id"`
	PersonID    string     `json:"person_id"`
	Testimony   string     `json:"testimony"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type CreateTestimony struct {
	ID          *int    `json:"id"`
	TestimonyID *string `json:"testimony_id"`
	PersonID    string  `json:"person_id"`
	Testimony   string  `json:"testimony"`
	Status      string  `json:"status"`
}

type UpdatePerson struct {
	ID          int     `json:"id"`
	PersonID    string  `json:"person_id"`
	Name        string  `json:"name"`
	Address     *string `json:"address"`
	PhoneNumber *string `json:"phone_number"`
	ClassName   string  `json:"class_name"`
}

type UpdateTestimony struct {
	ID          int    `json:"id"`
	TestimonyID string `json:"testimony_id"`
	PersonID    string `json:"person_id"`
	Testimony   string `json:"testimony"`
	Status      string `json:"status"`
}

type TestimonyStatus string

const (
	TestimonyStatusDraft  TestimonyStatus = "DRAFT"
	TestimonyStatusPublic TestimonyStatus = "PUBLIC"
	TestimonyStatusOther  TestimonyStatus = "OTHER"
)

var AllTestimonyStatus = []TestimonyStatus{
	TestimonyStatusDraft,
	TestimonyStatusPublic,
	TestimonyStatusOther,
}

func (e TestimonyStatus) IsValid() bool {
	switch e {
	case TestimonyStatusDraft, TestimonyStatusPublic, TestimonyStatusOther:
		return true
	}
	return false
}

func (e TestimonyStatus) String() string {
	return string(e)
}

func (e *TestimonyStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TestimonyStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TestimonyStatus", str)
	}
	return nil
}

func (e TestimonyStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
