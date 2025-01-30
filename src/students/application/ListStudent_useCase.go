package application

import (
	"ApiRestAct1/src/students/domain"
	"ApiRestAct1/src/students/domain/entities"
)

type ListStudent struct {
	db domain.IStudent
}

func NewListStudent(db domain.IStudent) *ListStudent {
	return &ListStudent{db: db}
}

func (ls *ListStudent) Execute() ([]entities.Student, error) {
	return ls.db.GetAll()
}
