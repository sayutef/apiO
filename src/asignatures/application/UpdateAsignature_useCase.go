package application

import (
	"ApiRestAct1/src/asignatures/domain"
	"ApiRestAct1/src/asignatures/domain/entities"
)

type UpdateAsignature struct {
	db domain.IAsignature
}

func NewUpdateAsignature(db domain.IAsignature) *UpdateAsignature {
	return &UpdateAsignature{db: db}
}

func (ua *UpdateAsignature) Execute(asignature entities.Asignature) error {
	return ua.db.Update(asignature)
}
