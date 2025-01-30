package application

import (
	"ApiRestAct1/src/asignatures/domain"
	"ApiRestAct1/src/asignatures/domain/entities"
)

type ListAsignature struct {
	db domain.IAsignature
}

func NewListAsignature(db domain.IAsignature) *ListAsignature {
	return &ListAsignature{db: db}
}

func (la *ListAsignature) Execute() ([]entities.Asignature, error) {
	return la.db.GetAll()
}
