package repo

import (
	"context"
	"orgService/internal/model"

	"gorm.io/gorm"
)

func (r *repo) ReassignDepartment(ctx context.Context, tx *gorm.DB, src, dst int) error {
	err := tx.WithContext(ctx).Model(&model.Employee{}).Where("department_id = ?", src).Update("department_id", dst).Error
	if err != nil {
		return err
	}

	err = tx.WithContext(ctx).Model(&model.Department{}).Where("parent_id = ?", src).Update("parent_id", dst).Error
	if err != nil {
		return err
	}

	return nil
}
