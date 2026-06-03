package repo

import (
	"context"
	"orgService/internal/model"

	"gorm.io/gorm"
)

func (r *repo) ReassignDepartment(ctx context.Context, tx *gorm.DB, src, dst int) error {
	res := tx.WithContext(ctx).Model(&model.Employee{}).Where("department_id = ?", src).Update("department_id", dst)
	if res.Error != nil {
		return res.Error
	}

	res = tx.WithContext(ctx).Model(&model.Department{}).Where("parent_id = ?", src).Update("parent_id", dst)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
