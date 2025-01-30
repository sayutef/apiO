package application

import (
	"ApiRestAct1/src/students/domain"
	"ApiRestAct1/src/students/domain/entities"
)

type ListStudentById struct {
	db domain.IStudent
}

func NewListStudentById(db domain.IStudent) *ListStudentById {
	return &ListStudentById{db: db}
}

func (ls_id *ListStudentById) Execute(id int) (entities.Student, error) {
	return ls_id.db.GetById(id)
}
