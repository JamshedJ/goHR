package models

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/JamshedJ/goHR/internal/configs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrNoRows       = errors.New("no rows in result set")
	ErrDuplicate    = errors.New("duplicate")
	ErrUnauthorized = errors.New("unauthorized")

	OK           = map[string]string{"message": "success"}
	NotFound     = map[string]string{"message": "not found"}
	Duplicate    = map[string]string{"message": "duplicate"}
	BadRequest   = map[string]string{"message": "bad request"}
	InternalErr  = map[string]string{"message": "internal server error"}
	Unauthorized = map[string]string{"message": "unauthorized"}

	roleAdmin = "admin"
)

type ErrorBadRequest struct {
	Message string `json:"message"`
}

func (e ErrorBadRequest) Error() string {
	return e.Message
}

func NewErrorBadRequest(msg string) ErrorBadRequest {
	return ErrorBadRequest{msg}
}

type Employee struct {
	ID             int     `json:"id,omitempty"`
	FirstName      string  `json:"first_name,omitempty"`
	LastName       string  `json:"last_name,omitempty"`
	PositionID     int     `json:"position_id,omitempty"`
	DepartmentID   int     `json:"department_id,omitempty"`
	EmploymentDate string  `json:"employment_date,omitempty"`
	Salary         float64 `json:"salary,omitempty"`
}

func (e *Employee) Validate() error {
	if len(e.FirstName) < 3 || len(e.FirstName) > 255 {
		return NewErrorBadRequest("invalid first_name")
	}
	if len(e.LastName) < 3 || len(e.LastName) > 255 {
		return NewErrorBadRequest("invalid last_name")
	}
	if e.PositionID <= 0 {
		return NewErrorBadRequest("invalid position_id")
	}
	if e.DepartmentID <= 0 {
		return NewErrorBadRequest("invalid department_id")
	}
	if _, err := time.Parse("2006-01-02", e.EmploymentDate); err != nil {
		return NewErrorBadRequest("invalid employment_date")
	}
	if e.Salary < 0 {
		return NewErrorBadRequest("invalid salary")
	}
	return nil
}

type Department struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	DepartmentHead string `json:"department_head"`
}

func (d *Department) Validate() error {
	if d.ID < 0 {
		return NewErrorBadRequest("invalid id")
	}
	if len(d.Title) > 255 {
		return NewErrorBadRequest("invalid title")
	}
	if len(d.DepartmentHead) > 255 {
		return NewErrorBadRequest("invalid department_head")
	}
	return nil
}

type Position struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Salary        float64 `json:"salary"`
	Qualification string  `json:"qualification"`
}

func (p *Position) Validate() error {
	if len(p.Title) > 255 {
		return NewErrorBadRequest("invalid title")
	}
	if len(p.Qualification) > 255 {
		return NewErrorBadRequest("invalid qualification")
	}
	if p.Salary < 0 {
		return NewErrorBadRequest("invalid salary")
	}
	return nil
}

type EmployeeRequest struct {
	ID                    int    `json:"id"`
	EmployeeID            int    `json:"employee_id"`
	StartsAt              string `json:"starts_at"`
	EndsAt                string `json:"ends_at"`
	Reason                string `json:"reason"`
	EmployeeRequestTypeID int    `json:"employee_request_type_id"`
}

func (e *EmployeeRequest) Validate() error {
	if e.ID < 0 {
		return NewErrorBadRequest("invalid id")
	}
	if e.EmployeeID <= 0 {
		return NewErrorBadRequest("invalid employee_id")
	}
	if _, err := time.Parse("2006-01-02", e.StartsAt); err != nil {
		return NewErrorBadRequest("invalid starts_at")
	}
	if _, err := time.Parse("2006-01-02", e.EndsAt); err != nil {
		return NewErrorBadRequest("invalid ends_at")
	}
	if len(e.Reason) > 255 {
		return NewErrorBadRequest("invalid_reason")
	}
	if e.EmployeeRequestTypeID <= 0 {
		return NewErrorBadRequest("invalid employee_request_type_id")
	}
	return nil
}

type EmployeeRequestType struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func (e *EmployeeRequestType) Validate() error {
	if e.ID < 0 {
		return NewErrorBadRequest("invalid id")
	}
	if len(e.Title) > 255 {
		return NewErrorBadRequest("invalid title")
	}
	return nil
}

type User struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role,omitempty"`
}

func (u *User) Validate() error {
	if len(u.Username) < 3 || len(u.Username) > 256 {
		return NewErrorBadRequest("invalid username")
	}
	if len(u.Password) < 4 || len(u.Password) > 256 {
		return NewErrorBadRequest("invalid password")
	}
	if u.Role != "" {
		if u.Role != "user" && u.Role != "admin" {
			return NewErrorBadRequest("invalid user role")
		}
	}
	return nil
}

func (u *User) IsAdmin() bool {
	return u.Role == roleAdmin
}

type JWTClaims struct {
	jwt.RegisteredClaims
	User User
}

func GeneratePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(configs.App.Salt)))
}
