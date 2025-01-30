package infraestructure

import (
	"ApiRestAct1/src/asignatures/application"
	"ApiRestAct1/src/asignatures/infraestructure/controllers"
	"ApiRestAct1/src/asignatures/infraestructure/database"
)

type DependenciesAsignature struct {
	CreateAsignatureController   *controllers.CreateAsignatureController
	ListAsignatureController     *controllers.ListAsignatureController
	ListAsignatureByIDController *controllers.ListAsignatureByIdController
	UpdateAsignatureController   *controllers.UpdateAsignatureController
	DeleteAsignatureController   *controllers.DeleteAsignatureController
}

func InitAsignature() *DependenciesAsignature {
	ps := database.NewMySQL()

	return &DependenciesAsignature{
		CreateAsignatureController:   controllers.NewCreateAsignatureController(application.NewCreateAsignature(ps)),
		ListAsignatureController:     controllers.NewListAsignatureController(application.NewListAsignature(ps)),
		ListAsignatureByIDController: controllers.NewListAsignatureByIdController(application.NewListAsignatureById(ps)),
		UpdateAsignatureController:   controllers.NewUpdateAsignatureController(application.NewUpdateAsignature(ps)),
		DeleteAsignatureController:   controllers.NewDeleteAsignatureController(application.NewDeleteAsignature(ps)),
	}
}
