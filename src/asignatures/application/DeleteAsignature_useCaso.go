package application

import (
	"ApiRestAct1/src/asignatures/domain"
)

type DeleteAsignature struct {
	db domain.IAsignature
}

func NewDeleteAsignature(db domain.IAsignature) *DeleteAsignature {
	return &DeleteAsignature{db: db}
}

func (da *DeleteAsignature) Execute(id int) error {
	return da.db.Delete(id)
}
