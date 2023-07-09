package repository

import (
	"context"

	"github.com/faba/client/domain/model"
)

// FabaRepository implements interface to request
type FabaRepository interface {
	CreateEmployee(ctx context.Context, employeeRequestParam *model.EmployeeRequestParam) (model.EmployeeResponseData, error)
}
