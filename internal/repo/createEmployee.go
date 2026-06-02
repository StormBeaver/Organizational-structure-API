package repo

import (
	"context"
	"orgService/internal/model"
	"time"
)

func (r *repo) CreateEmployee(ctx context.Context, name string, position string, depID int, hiredAt *time.Time) (model.Employee, error) {
	employee := model.Employee{
		DepartmentID: depID,
		FullName:     name,
		Position:     position,
		HiredAt:      hiredAt,
	}
	res := r.db.WithContext(ctx).Model(model.Employee{}).Create(employee) // проверить работоспособность т.к. на сайте db.Select("Name", "Age", "CreatedAt").Create(&user)

	if res.Error != nil {
		return employee, res.Error
	}

	return employee, nil
}
