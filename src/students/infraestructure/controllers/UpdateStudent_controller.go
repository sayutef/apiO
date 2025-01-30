package controllers

import (
	"ApiRestAct1/src/students/application"
	"ApiRestAct1/src/students/domain/entities"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UpdateStudentController struct {
	useCase *application.UpdateStudent
}

func NewUpdateStudentController(useCase *application.UpdateStudent) *UpdateStudentController {
	return &UpdateStudentController{useCase: useCase}
}

func (us_c *UpdateStudentController) Execute(c *gin.Context) {
	var student entities.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID no válido"})
		return
	}
	student.ID = idInt
	if err := us_c.useCase.Execute(student); err != nil {
		c.JSON(500, gin.H{"error": "No se pudo actualizar el estudiante"})
		return
	}
	c.JSON(200, gin.H{"message": "Estudiante actualizado"})
}
