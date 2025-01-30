package application

import (
	"ApiRestAct1/src/asignatures/domain"
	"ApiRestAct1/src/asignatures/domain/entities"
)

type CreateAsignature struct {
	db domain.IAsignature
}

func NewCreateAsignature(db domain.IAsignature) *CreateAsignature {
	return &CreateAsignature{db: db}
}

func (ca *CreateAsignature) Execute(asignature entities.Asignature) error {
	return ca.db.Save(asignature)
}
