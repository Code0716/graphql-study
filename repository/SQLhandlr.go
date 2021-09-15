package repository

import "github.com/Code0716/graphql-study/domain"

// SQLHandlerInterface  SQLHandler
type SQLHandlerInterface interface {
	Create(value interface{}) error
	Find(value interface{}, pager domain.Pager) error
	First(value interface{}, where ...interface{}) error
}
