package application

import "ApiRestAct1/src/students/domain"

type DeleteStudent struct {
	db domain.IStudent
}

func NewDeleteStudent(db domain.IStudent) *DeleteStudent {
	return &DeleteStudent{db: db}
}

func (ds *DeleteStudent) Execute(id int) error {
	return ds.db.Delete(id)
}
