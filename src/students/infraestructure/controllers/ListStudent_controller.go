package controllers

import (
	"ApiRestAct1/src/students/application"
	"github.com/gin-gonic/gin"
)

type ListStudentController struct {
	useCase *application.ListStudent
}

func NewListStudentController(useCase *application.ListStudent) *ListStudentController {
	return &ListStudentController{useCase: useCase}
}

func (ls_c *ListStudentController) Execute(c *gin.Context) {
	students, err := ls_c.useCase.Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener los estudiantes"})
		return
	}
	c.JSON(200, students)
}
