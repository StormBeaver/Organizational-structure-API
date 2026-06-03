package repo

import (
	"context"
	"fmt"
	"orgService/internal/model"

	"gorm.io/gorm"
)

func (r *repo) DeleteDepartment(ctx context.Context, tx *gorm.DB, department *model.Department) error {
	res := tx.WithContext(ctx).Delete(department)

	if res.Error != nil {
		return fmt.Errorf("gorm delete department: %w", res.Error)
	}

	return nil
}
