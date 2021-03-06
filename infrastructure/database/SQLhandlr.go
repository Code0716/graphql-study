package database

import (
	"errors"

	"github.com/Code0716/graphql-study/domain"
	"gorm.io/gorm"
)

// SQLHandler struct
type SQLHandler struct {
	Conn *gorm.DB
}

// Create retuern error
func (h SQLHandler) Create(value interface{}) error {
	err := h.Conn.Create(value).Error
	if err != nil {
		return err
	}
	return nil
}

// Find gorm find
func (h SQLHandler) Find(value interface{}, pager domain.Pager, where ...interface{}) error {
	err := h.Conn.
		Limit(*pager.Limit).
		Offset(*pager.Offset).
		Find(value, where...).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	if err != nil {
		return err
	}
	return nil
}

// First gorm find
func (h SQLHandler) First(value interface{}, where ...interface{}) error {
	err := h.Conn.First(value, where...).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	if err != nil {
		return err
	}
	return nil
}

// Save is gorm save
// TODO:未実装
func (h *SQLHandler) Save(value interface{}) error {
	err := h.Conn.Save(value).Error
	return err
}

// Update is gorm Update
// TODO:未実装
func (h *SQLHandler) Update(column string, value interface{}) error {
	err := h.Conn.Update(column, value).Error
	return err
}

// Scan is gorm Scan
// TODO 実装未
func (h *SQLHandler) Scan(value interface{}) error {
	err := h.Conn.Scan(value).Error
	return err
}

// Exec is gorm Exec
// TODO 実装未
func (h *SQLHandler) Exec(sql string, value ...interface{}) error {
	err := h.Conn.Exec(sql, value).Error
	return err
}

// Raw is gorm Raw
// TODO 実装未
func (h *SQLHandler) Raw(sql string, values ...interface{}) *gorm.DB {
	return h.Conn.Raw(sql, values...)
}

// Delete is gorm delete
// TODO 実装未
func (h *SQLHandler) Delete(value interface{}, where ...interface{}) error {
	err := h.Conn.Delete(value, where...).Error
	return err
}

// Where gorm Where
func (h *SQLHandler) Where(query interface{}, args ...interface{}) error {
	err := h.Conn.Where(query, args...).Error
	return err
}

func (h SQLHandler) IsExist(tableName string, query interface{}, args ...interface{}) (bool, error) {
	var count int64

	err := h.Conn.Table(tableName).
		Where(query, args...).
		Count(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}
