package service

import (
	"context"

	"github.com/JamshedJ/goHR/internal/log"

	"github.com/JamshedJ/goHR/internal/models"
)

func (s *Service) CreateEmployee(ctx context.Context, e models.Employee) (id int, err error) {
	if !e.Validate() {
		return 0, models.ErrBadRequest
	}
	id, err = s.db.CreateEmployee(ctx, e)
	if err != nil {
		log.Error.Println("app CreateEmployee", err)
	}
	return
}

func (s *Service) GetEmployeeByID(ctx context.Context, id int, isAdmin bool) (employee models.Employee, err error) {
	if id <= 0 {
		return employee, models.ErrBadRequest
	}
	employee, err = s.db.GetEmployeeByID(ctx, id)
	if err != nil {
		if err != models.ErrNoRows {
			log.Error.Println("app GetEmployeeByID", err)
		}
		return
	}
	if !isAdmin {
		employee.Salary = 0
	}
	return
}

func (s *Service) GetAllEmployees(ctx context.Context, isAdmin bool) (employees []models.Employee, err error) {
	employees, err = s.db.GetAllEmployees(ctx)
	if err != nil {
		if err != models.ErrNoRows {
			log.Error.Println("app GetEmployees", err)
		}
		return
	}
	if !isAdmin {
		for _, e := range employees {
			e.Salary = 0
		}
	}
	return
}

func (s *Service) UpdateEmployee(ctx context.Context, e models.Employee) (err error) {
	if !e.Validate() || e.ID <= 0 {
		return models.ErrBadRequest
	}

	err = s.db.UpdateEmployee(ctx, e)
	if err != nil {
		log.Error.Println("app UpdateEmployee", err)
	}
	return
}

func (s *Service) DeleteEmployee(ctx context.Context, id int) (err error) {
	if id <= 0 {
		return models.ErrBadRequest
	}

	err = s.db.DeleteEmployee(ctx, id)
	if err != nil {
		log.Error.Println("app DeleteEmployee", err)
	}
	return
}

func (s *Service) SearchEmployeeByName(ctx context.Context, query string) (employees []models.Employee, err error) {
	employees, err = s.db.SearchEmployeeByName(ctx, query)
	if err != nil {
		log.Error.Println("app SearchEmployeeByName", err)
	}
	return
}
