package infraestructure

import (
	"ApiRestAct1/src/asignatures/application"
	"ApiRestAct1/src/asignatures/infraestructure/adapter"
	"ApiRestAct1/src/asignatures/infraestructure/controllers"
	"ApiRestAct1/src/asignatures/infraestructure/database"
	"log"
)

type DependenciesAsignature struct {
	CreateAsignatureController   *controllers.CreateAsignatureController
	ListAsignatureController     *controllers.ListAsignatureController
	ListAsignatureByIDController *controllers.ListAsignatureByIdController
	UpdateAsignatureController   *controllers.UpdateAsignatureController
	DeleteAsignatureController   *controllers.DeleteAsignatureController
}

// In infraestructure/dependencies.go
func InitAsignature() *DependenciesAsignature {
	ps := database.NewMySQL()

	// Initialize RabbitMQ client
	rmqClient, err := adapter.NewRabbitMQAdapter()
	if err != nil {
		log.Fatalf("Error initializing RabbitMQ: %v", err)
	}

	// Pass both application.CreateAsignature, ps, and rmqClient to the controller constructor
	return &DependenciesAsignature{
		CreateAsignatureController:   controllers.NewCreateAsignatureController(application.NewCreateAsignature(ps, rmqClient), ps), // Pass both parameters
		ListAsignatureController:     controllers.NewListAsignatureController(application.NewListAsignature(ps)),
		ListAsignatureByIDController: controllers.NewListAsignatureByIdController(application.NewListAsignatureById(ps)),
		UpdateAsignatureController:   controllers.NewUpdateAsignatureController(application.NewUpdateAsignature(ps)),
		DeleteAsignatureController:   controllers.NewDeleteAsignatureController(application.NewDeleteAsignature(ps)),
	}
}
