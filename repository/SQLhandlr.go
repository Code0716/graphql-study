package repository

import "github.com/Code0716/graphql-study/domain"

// SQLHandlerInterface  SQLHandler
type SQLHandlerInterface interface {
	Create(value interface{}) error
	Find(value interface{}, pager domain.Pager, where ...interface{}) error
	First(value interface{}, where ...interface{}) error
	IsExist(tableName string, query interface{}, args ...interface{}) (bool, error)
}
