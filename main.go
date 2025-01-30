package main

import (
	"ApiRestAct1/src/students/infraestructure"
	"ApiRestAct1/src/students/infraestructure/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	dependencies := infraestructure.InitStudent()

	r := gin.Default()

	routes.ConfigureRoutes(r,
		dependencies.CreateStudentController,
		dependencies.ListStudentController,
		dependencies.ListStudentByIDController,
		dependencies.UpdateStudentController,
		dependencies.DeleteStudentController)

	r.Run(":8080")
}
