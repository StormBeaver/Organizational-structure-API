package repo

import (
	"context"
	"orgService/internal/model"
)

func (r *repo) PatchDepartment(ctx context.Context, name string, ID int, parentID *int) (model.Department, error) {
	department := model.Department{}
	res := r.db.WithContext(ctx).Model(department).Where("id = ? AND name = ?", ID, name).Update("parent_id", parentID)
	return department, res.Error
}
