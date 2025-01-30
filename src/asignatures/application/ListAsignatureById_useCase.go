package application

import (
	"ApiRestAct1/src/asignatures/domain"
	"ApiRestAct1/src/asignatures/domain/entities"
)

type ListAsignatureById struct {
	db domain.IAsignature
}

func NewListAsignatureById(db domain.IAsignature) *ListAsignatureById {
	return &ListAsignatureById{db: db}
}

func (la_id *ListAsignatureById) Execute(id int) (entities.Asignature, error) {
	return la_id.db.GetById(id)
}
