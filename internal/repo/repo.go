package repo

import (
	"context"
	"orgService/internal/model"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo {
	return &repo{db: db}
}

func (r *repo) LastDepartment(ctx context.Context) (model.Department, error) {
	department := model.Department{}

	res := r.db.WithContext(ctx).Model(model.Department{}).Last(&department)
	if res.Error != nil {
		return department, res.Error
	}
	return department, nil
}
