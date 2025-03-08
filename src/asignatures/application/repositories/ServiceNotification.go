package repositories

import (
	"ApiRestAct1/src/asignatures/domain/entities"
	"log"
)

type ServiceNotification struct {
	imageService IMessageService
}

func NewServiceNotification(imageService IMessageService) *ServiceNotification {
	return &ServiceNotification{imageService: imageService}
}

func (sn *ServiceNotification) NotifyAsignatureCreated(asignature entities.Asignature) error {
	log.Println("Notificando la creacion de las asignaturas ...")

	err := sn.imageService.PublishEvent("AppointmentCreated", asignature)
	if err != nil {
		log.Println("Error al publicar el evento %v", err)
		return err
	}
	return nil
}
