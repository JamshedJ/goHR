package models

import (
	"errors"
)

type Employee struct {
	ID             int     `json:"id"`
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	Position       string  `json:"position"`
	Department     string  `json:"department"`
	EmploymentDate string  `json:"employment_date"`
	Salary         float64 `json:"salary"`
}

type Department struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	DepartmentHead string `json:"department_head"`
}

type Position struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Salary        float64 `json:"salary"`
	Qualification string  `json:"qualification"`
}

type VacationRequest struct {
	ID         int    `json:"id"`
	EmployeeID int    `json:"employee_id"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	Reason     string `json:"reason"`
}

type SickLeaveRequest struct {
	ID         int    `json:"id"`
	EmployeeID int    `json:"employee_id"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	Reason     string `json:"reason"`
}

type User struct {
	ID       uint   `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role"`
}

var ErrNoRows = errors.New("no rows in result set")
var ErrDuplicate = errors.New("duplicate")
var ErrBadRequest = errors.New("bad request")
var ErrUnauthorized = errors.New("unauthorized")

var (
	OK           = map[string]string{"message": "success"}
	NotFound     = map[string]string{"message": "not found"}
	Duplicate    = map[string]string{"message": "duplicate"}
	BadRequest   = map[string]string{"message": "bad request"}
	InternalErr  = map[string]string{"message": "internal server error"}
	Unauthorized = map[string]string{"message": "unauthorized"}
)
