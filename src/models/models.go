package models

import (
	"errors"
	"time"
)

// IModel Base model interface
type IModel interface {
	IsNewRecord() bool
	Find(id int64) (*IModel, error)
	Save() (*IModel, error)
	Destroy() (*IModel, error)
}

// ICollection Base collection interface
type ICollection interface {
	All() (*[]IModel, error)
}

// Model Parent struct of all records
type Model struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ModelsCollection Collection of models
type ModelsCollection []Model

// All Returns all records from database
func (records ModelsCollection) All() (*ModelsCollection, error) {
	return nil, errors.New("Func is not implemented")
}

// Find Return record from database
func (model *Model) Find(id int) (*Model, error) {
	return model, errors.New("Func is not implemented")
}

// Save Save record in database
func (model *Model) Save() (*Model, error) {
	return nil, errors.New("Func is not implemented")
}

// Destroy Delete record from database
func (model *Model) Destroy() (*Model, error) {
	return nil, errors.New("Func is not implemented")
}

// IsNewRecord Returns true in record can't be saved in database
func (model *Model) IsNewRecord() bool {
	return model.ID == 0
}
