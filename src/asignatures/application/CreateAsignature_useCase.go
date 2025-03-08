package application

import (
	"ApiRestAct1/src/asignatures/application/repositories"
	"ApiRestAct1/src/asignatures/domain"
	"ApiRestAct1/src/asignatures/domain/entities"
	"log"
)

type CreateAsignature struct {
	asignatureRepo      domain.IAsignature
	serviceNotification repositories.IMessageService
}

func NewCreateAsignature(asignatureRepo domain.IAsignature, serviceNotification repositories.IMessageService) *CreateAsignature {
	return &CreateAsignature{
		asignatureRepo:      asignatureRepo,
		serviceNotification: serviceNotification,
	}
}

func (c *CreateAsignature) Execute(asignature entities.Asignature) (entities.Asignature, error) {
	created, err := c.asignatureRepo.Save(asignature)
	if err != nil {
		return entities.Asignature{}, err
	}

	err = c.serviceNotification.PublishEvent("AsignatureCreated", created)
	if err != nil {
		log.Printf("Error notifying about created asignature: %v", err)
		return entities.Asignature{}, err
	}

	return created, nil
}
