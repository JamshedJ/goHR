package models

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Employee struct {
	ID             int     `json:"id"`
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	PositionID     int     `json:"position_id"`
	DepartmentID   int     `json:"department_id"`
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
	Role     string `json:"role,omitempty"`
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

func ReplyError(c *gin.Context, err error) {
	switch err {
	case ErrUnauthorized:
		c.JSON(http.StatusUnauthorized, Unauthorized)
	case ErrBadRequest:
		c.JSON(http.StatusBadRequest, BadRequest)
	case ErrNoRows:
		c.JSON(http.StatusNotFound, NotFound)
	case ErrDuplicate:
		c.JSON(http.StatusNotAcceptable, Duplicate)
	default:
		c.JSON(http.StatusInternalServerError, InternalErr)
	}
	return
}

func (e *Employee) Validate() bool {
	if e.ID <= 0 {
		return false
	}
	if len(e.FirstName) < 3 || len(e.FirstName) > 128 {
		return false
	}
	if len(e.LastName) < 3 || len(e.LastName) > 128 {
		return false
	}
	return true
}

func (u *User) Validate() bool {
	if len(u.Username) < 3 || len(u.Username) > 128 {
		return false
	}
	if len(u.Password) < 4 || len(u.Password) > 32 {
		return false
	}
	return true
}

type JWTClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}
