package repo

import (
	"context"
	"fmt"
	"orgService/internal/model"
)

func (r *repo) CreateDepartment(ctx context.Context, department *model.Department) (*model.Department, error) {
	err := r.db.WithContext(ctx).Model(model.Department{}).Create(&department).Error

	if err != nil {
		return nil, fmt.Errorf("gorm create department: %w", err)
	}

	return department, nil
}
