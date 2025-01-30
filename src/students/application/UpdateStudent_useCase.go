package application

import (
	"ApiRestAct1/src/students/domain"
	"ApiRestAct1/src/students/domain/entities"
)

type UpdateStudent struct {
	db domain.IStudent
}

func NewUpdateStudent(db domain.IStudent) *UpdateStudent {
	return &UpdateStudent{db: db}
}

func (us *UpdateStudent) Execute(student entities.Student) error {
	return us.db.Update(student)
}
