package models

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrNoRows       = errors.New("no rows in result set")
	ErrDuplicate    = errors.New("duplicate")
	ErrBadRequest   = errors.New("bad request")
	ErrUnauthorized = errors.New("unauthorized")

	OK           = map[string]string{"message": "success"}
	NotFound     = map[string]string{"message": "not found"}
	Duplicate    = map[string]string{"message": "duplicate"}
	BadRequest   = map[string]string{"message": "bad request"}
	InternalErr  = map[string]string{"message": "internal server error"}
	Unauthorized = map[string]string{"message": "unauthorized"}

	RoleAdmin = "admin"
	RoleUser  = "user"
)

type Employee struct {
	ID             int     `json:"id,omitempty"`
	FirstName      string  `json:"first_name,omitempty"`
	LastName       string  `json:"last_name,omitempty"`
	PositionID     int     `json:"position_id,omitempty"`
	DepartmentID   int     `json:"department_id,omitempty"`
	EmploymentDate string  `json:"employment_date,omitempty"`
	Salary         float64 `json:"salary,omitempty"`
}

func (e *Employee) Validate() bool {
	if len(e.FirstName) < 3 || len(e.FirstName) > 255 {
		return false
	}
	if len(e.LastName) < 3 || len(e.LastName) > 255 {
		return false
	}
	if e.PositionID <= 0 {
		return false
	}
	if e.DepartmentID <= 0 {
		return false
	}
	if _, err := time.Parse("2006-01-02", e.EmploymentDate); err != nil {
		return false
	}
	if e.Salary < 0 {
		return false
	}
	return true
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

type Request struct {
	ID         int    `json:"id"`
	EmployeeID int    `json:"employee_id"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	Reason     string `json:"reason"`
	TypeID     int    `json:"type_id"`
}

func (r *Request) Validate() bool {
	if _, err := time.Parse("2006-01-02", r.StartDate); err != nil {
		return false
	}
	if _, err := time.Parse("2006-01-02", r.EndDate); err != nil {
		return false
	}
	return true
}

type Type struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type User struct {
	ID       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role,omitempty"`
}

func (u *User) Validate() bool {
	if len(u.Username) < 3 || len(u.Username) > 256 {
		return false
	}
	if len(u.Password) < 4 || len(u.Password) > 256 {
		return false
	}
	if u.Role != "" {
		if u.Role != "user" && u.Role != "admin" {
			return false
		}
	}
	return true
}

func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

type JWTClaims struct {
	jwt.RegisteredClaims
	User User
}
