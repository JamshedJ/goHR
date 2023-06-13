package models

import "gorm.io/gorm"

type Employee struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Email        string  `json:"email"`
	Address      string  `json:"address"`
	Salary       float64 `json:"salary"`
	Position     string  `json:"position"`
	VocationDay  int     `json:"vocation_day"`
	DepartmentID uint    `json:"department_id"`
	Department   Department
}

type Department struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	DepartmentName string `json:"department_name"`
}

type Vacancy struct {
	gorm.Model
	Title        string `json:"title" gorm:"size:128;not null"`
	Description  string `json:"description" gorm:"size:256"`
	Requirements string `json:"requirements" gorm:"size:256"`
}

type Users struct {
	ID       uint   `json:"id,omitempty" gorm:"primaryKey"`
	Username string `json:"username,omitempty" gorm:"unique;size:32;not null"`
	Password string `json:"password,omitempty" gorm:"size:128;not null;->:false;<-"`
}
