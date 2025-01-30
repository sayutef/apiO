package routes

import (
	"ApiRestAct1/src/students/infraestructure/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(r *gin.Engine,
	createStudentController *controllers.CreateStudentController,
	listStudentController *controllers.ListStudentController,
	listStudentByIDController *controllers.ListStudentByIDController,
	updateStudentController *controllers.UpdateStudentController,
	deleteStudentController *controllers.DeleteStudentController) {

	r.GET("/students", listStudentController.Execute)
	r.GET("/students/:id", listStudentByIDController.Execute)
	r.POST("/students", createStudentController.Execute)
	r.PUT("/students/:id", updateStudentController.Execute)
	r.DELETE("/students/:id", deleteStudentController.Execute)
}
