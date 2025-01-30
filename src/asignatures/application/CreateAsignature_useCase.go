package application

import "ApiRestAct1/src/asignatures/domain"

type CreateAsignature struct {
	db domain.IAsignature
}

func NewCreateAsignature(db domain.IAsignature) CreateAsignature {
	return CreateAsignature{db: db}
}


func (ca *CreateAsignature) Execute() {
	ca.db.Save()
}
