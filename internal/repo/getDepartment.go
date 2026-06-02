package repo

import (
	"context"
	"orgService/internal/model"
)

// я хз как это сделать...
func (r *repo) GetDepartment(ctx context.Context, depth int, employees bool) (model.Department, *[]model.Employee, error) {
	res := r.db.WithContext(ctx).Model(model.Department{}).Where("").Find(&model.Department{})
	return model.Department{}, nil, res.Error
}
