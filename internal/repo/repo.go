package repo

import (
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo {
	return &repo{db: db}
}

func (r *repo) BeginTx() *gorm.DB {
	return r.db.Begin()
}
