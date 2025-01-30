package controllers

import (
	"ApiRestAct1/src/students/application"
	"ApiRestAct1/src/students/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateStudentController struct {
	useCase *application.CreateStudent
}

func NewCreateStudentController(useCase *application.CreateStudent) *CreateStudentController {
	return &CreateStudentController{useCase: useCase}
}

func (cs_c *CreateStudentController) Execute(c *gin.Context) {
	var student entities.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}
	if err := cs_c.useCase.Execute(student); err != nil {
		c.JSON(500, gin.H{"error": "No se pudo crear el estudiante"})
		return
	}
	c.JSON(201, gin.H{"message": "Estudiante creado correctamente"})
}
