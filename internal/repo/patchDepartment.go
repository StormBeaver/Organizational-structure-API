package repo

import (
	"context"
	"fmt"
	"orgService/internal/model"
)

func (r *repo) PatchDepartment(ctx context.Context, department *model.Department) (*model.Department, error) {
	err := r.db.WithContext(ctx).Save(department).Error

	if err != nil {
		return nil, fmt.Errorf("gorm patch department: %w", err)
	}
	return department, nil
}
