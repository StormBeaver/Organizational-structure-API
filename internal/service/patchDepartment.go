package service

import (
	"context"
	"fmt"
	appErrors "orgService/internal/errors"
	"orgService/internal/handlers/dto"
	"orgService/internal/model"
)

func (s *Service) PatchDepartment(ctx context.Context, request dto.PatchDepartmentRequest) (*model.Department, error) {

	if request.ParentID == nil && request.Name == nil {
		s.logger.Err(appErrors.ErrInvalidArguments).Msg("invalid arguments")
		return nil, fmt.Errorf("no arguments: %w", appErrors.ErrInvalidArguments)
	}

	dep, err := s.repo.GetDepartment(ctx, request.Id)
	if err != nil {
		s.logger.Err(appErrors.ErrInvalidDepartmentNumber).Msg("department doesn't exist")
		return nil, fmt.Errorf("department doesn't exist: %w", appErrors.ErrInvalidDepartmentNumber)
	}

	if request.Name != nil {
		err := validateFieldLength(*request.Name)
		if err != nil {
			s.logger.Err(appErrors.ErrInvalidFieldLength).Msg("invalid name length")
			return nil, fmt.Errorf("name invalid: %w", appErrors.ErrInvalidFieldLength)
		}

		dep.Name = request.Name
	}

	if request.ParentID != nil {
		parent, err := s.repo.GetDepartment(ctx, *request.ParentID)
		if err != nil {
			s.logger.Err(appErrors.ErrInvalidDepartmentNumber).Msg("parent department doesn't exist")
			return nil, fmt.Errorf("parent department doesn't exist: %w", appErrors.ErrInvalidDepartmentNumber)
		}
		dep.Parent = parent
	}

	return s.repo.PatchDepartment(ctx, dep)
}

func (s *Service) validatePatchDepartment(ctx context.Context, request dto.PatchDepartmentRequest) error {

	return nil
}
