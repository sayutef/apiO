package routes

import (
	"ApiRestAct1/src/asignatures/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutesAsignature(
	r *gin.Engine,
	createAsignatureController *controllers.CreateAsignatureController,
	listAsignatureController *controllers.ListAsignatureController,
	listAsignatureByIDController *controllers.ListAsignatureByIdController,
	updateAsignatureController *controllers.UpdateAsignatureController,
	deleteAsignatureController *controllers.DeleteAsignatureController,
) {
	r.GET("/asignatures", listAsignatureController.Execute)
	r.GET("/asignatures/:id", listAsignatureByIDController.Execute)
	r.POST("/asignatures", createAsignatureController.Execute)
	r.PUT("/asignatures/:id", updateAsignatureController.Execute)
	r.DELETE("/asignatures/:id", deleteAsignatureController.Execute)
}
