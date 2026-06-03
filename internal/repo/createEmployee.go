package repo

import (
	"context"
	"fmt"
	"orgService/internal/model"
)

func (r *repo) CreateEmployee(ctx context.Context, employee *model.Employee) (*model.Employee, error) {
	err := r.db.WithContext(ctx).Model(model.Employee{}).Create(&employee).Error

	if err != nil {
		return nil, fmt.Errorf("gorm create employee: %w", err)
	}

	return employee, nil
}
