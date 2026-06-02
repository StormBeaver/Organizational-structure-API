package service

import (
	"context"
	"orgService/internal/model"
	"strings"
)

func (s *Service) CreateDepartment(ctx context.Context, name string, parentID *int) (model.Department, error) {
	err := s.validateCreateDepartment(ctx, name, parentID)
	if err != nil {
		return model.Department{}, err
	}

	name = strings.Trim(name, " ")

	return s.repo.CreateDepartment(ctx, name, parentID)
}
