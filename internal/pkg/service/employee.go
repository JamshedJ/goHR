package service

import (
	"context"

	"github.com/JamshedJ/goHR/internal/log"

	"github.com/JamshedJ/goHR/internal/models"
)

func (s *Service) CreateEmployee(ctx context.Context, e models.Employee) (id int, err error) {
	if err = e.Validate(); err != nil {
		log.Warning.Println("service CreateEmployee", err)
		return
	}
	id, err = s.db.CreateEmployee(ctx, e)
	if err != nil {
		log.Error.Println("service CreateEmployee", err)
	}
	return
}

func (s *Service) GetEmployeeByID(ctx context.Context, id int, isAdmin bool) (employee models.Employee, err error) {
	if id <= 0 {
		err = models.NewErrorBadRequest("invalid id")
		log.Warning.Println("service GetEmployeeByID", err)
		return
	}
	employee, err = s.db.GetEmployeeByID(ctx, id)
	if err != nil {
		if err == models.ErrNoRows {
			log.Warning.Println("service GetEmployeeByID", err)
			return
		}
		log.Error.Println("service GetEmployeeByID", err)
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
		if err == models.ErrNoRows {
			log.Warning.Println("service GetEmployee", err)
			return
		}
		log.Error.Println("service GetEmployees", err)
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
	if e.ID <= 0 {
		err = models.NewErrorBadRequest("invalid id")
	} else {
		err = e.Validate()
	}
	if err != nil {
		log.Warning.Println("service UpdateEmployee", err)
		return
	}

	err = s.db.UpdateEmployee(ctx, e)
	if err != nil {
		log.Error.Println("service UpdateEmployee", err)
	}
	return
}

func (s *Service) DeleteEmployee(ctx context.Context, id int) (err error) {
	if id <= 0 {
		err = models.NewErrorBadRequest("invalid id")
		log.Warning.Println("servie DeleteEmployee", err)
		return
	}

	err = s.db.DeleteEmployee(ctx, id)
	if err != nil {
		log.Error.Println("service DeleteEmployee", err)
	}
	return
}

func (s *Service) SearchEmployeeByName(ctx context.Context, query string) (employees []models.Employee, err error) {
	employees, err = s.db.SearchEmployeeByName(ctx, query)
	if err != nil {
		log.Error.Println("service SearchEmployeeByName", err)
	}
	return
}
