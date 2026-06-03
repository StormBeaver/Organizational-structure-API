package service

import (
	"context"
	"fmt"
	appErrors "orgService/internal/errors"
	"orgService/internal/handlers/dto"
)

func (s *Service) DeleteDepartment(ctx context.Context, request dto.DeleteDepartmentRequest) error {

	src, err := s.repo.GetDepartment(ctx, request.Id)
	if err != nil {
		return fmt.Errorf("deparment doesn't exist: %w", appErrors.ErrInvalidDepartmentNumber)
	}

	tx := s.repo.BeginTx()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if request.Mode == dto.ModeDeleteReassign {
		dst, err := s.repo.GetDepartment(ctx, *request.ToDepartment)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("deparment doesn't exist: %w", appErrors.ErrInvalidDepartmentNumber)
		}

		current := dst.Parent
		for current != nil {
			if current.Id == src.Id {
				tx.Rollback()
				return fmt.Errorf("department cycle detected: %w", appErrors.ErrInvalidDepartmentNumber)
			}

			if current.ParentId == nil {
				break
			}
			current, err = s.repo.GetDepartment(ctx, *current.ParentId)
			if err != nil {
				tx.Rollback()
				return err
			}

		}

		err = s.repo.ReassignDepartment(ctx, tx, src.Id, dst.Id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = s.repo.DeleteDepartment(ctx, tx, src)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
