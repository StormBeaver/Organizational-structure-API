package repo

import (
	"context"
	"fmt"
	"orgService/internal/model"

	"gorm.io/gorm"
)

func (r *repo) DeleteDepartment(ctx context.Context, tx *gorm.DB, department *model.Department) error {
	err := tx.WithContext(ctx).Delete(department).Error

	if err != nil {
		return fmt.Errorf("gorm delete department: %w", err)
	}

	return nil
}
