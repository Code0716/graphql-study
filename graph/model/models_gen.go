// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type CommonSuccessResponse struct {
	Message string `json:"Message"`
}

type CreateMember struct {
	Name        string     `json:"name"`
	Address     *string    `json:"address"`
	PhoneNumber *string    `json:"phone_number"`
	Status      *Status    `json:"status"`
	Birthday    *time.Time `json:"birthday"`
}

type GetMemberParams struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}

type Member struct {
	ID          int        `json:"id"`
	MemberID    string     `json:"memberId"`
	Name        string     `json:"name"`
	Address     string     `json:"address"`
	PhoneNumber string     `json:"phone_number"`
	Status      Status     `json:"status"`
	Birthday    *time.Time `json:"birthday"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
}

type Pager struct {
	Limit  *int    `json:"limit"`
	Offset *int    `json:"offset"`
	Status *string `json:"status"`
}

type Status string

const (
	StatusFriend   Status = "FRIEND"
	StatusFamily   Status = "FAMILY"
	StatusBusiness Status = "BUSINESS"
	StatusOther    Status = "OTHER"
)

var AllStatus = []Status{
	StatusFriend,
	StatusFamily,
	StatusBusiness,
	StatusOther,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusFriend, StatusFamily, StatusBusiness, StatusOther:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
