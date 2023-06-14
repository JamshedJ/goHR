package repository

type Employee interface {
}

type Department interface {
}

type Vacancy interface {
}

type User interface {
}

type Repository struct {
	Employee
	Department
	Vacancy
	User
}

func NewRepository(cfg string) *Repository {
	return &Repository{}
}
