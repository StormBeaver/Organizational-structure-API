package repo

import (
	"context"
	"fmt"
	"orgService/internal/model"
)

func (r *repo) PatchDepartment(ctx context.Context, department *model.Department) (*model.Department, error) {

	res := r.db.WithContext(ctx).Save(department)

	if res.Error != nil {
		return nil, fmt.Errorf("gorm patch department: %w", res.Error)
	}
	return department, nil
}
