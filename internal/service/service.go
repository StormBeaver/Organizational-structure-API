package service

import (
	"context"
	"orgService/internal/model"
	"time"

	"github.com/rs/zerolog"
)

type Repo interface {
	CreateDepartment(ctx context.Context, name string, parentID *int) (model.Department, error)
	CreateEmployee(ctx context.Context, name string, position string, depID int, hiredAt *time.Time) (model.Employee, error)
	GetDepartment(ctx context.Context, depth int, employees bool) (model.Department, *[]model.Employee, error)
	PatchDepartment(ctx context.Context, name string, ID int, parentID *int) (model.Department, error)
	DeleteDepartment(ctx context.Context, ID int, mode string, reasignDestination *int) error

	LastDepartment(ctx context.Context) (model.Department, error)
}

type Service struct {
	repo Repo

	logger *zerolog.Logger
}

func NewService(repo Repo, logger *zerolog.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) GetDepartment(ctx context.Context, depth int, employees bool) (model.Department, *[]model.Employee, error) {
	return s.repo.GetDepartment(ctx, depth, employees)
}

func (s *Service) PatchDepartment(ctx context.Context, name string, ID int, parentID *int) (model.Department, error) {
	return s.repo.PatchDepartment(ctx, name, ID, parentID)
}

func (s *Service) DeleteDepartment(ctx context.Context, ID int, mode string, reasignDestination *int) error {
	return s.repo.DeleteDepartment(ctx, ID, mode, reasignDestination)
}

func (s *Service) LastDepartment(ctx context.Context) (model.Department, error) {
	return s.repo.LastDepartment(ctx)
}
