package controllers

import (
	"ApiRestAct1/src/students/application"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ListStudentByIDController struct {
	useCase *application.ListStudentById
}

func NewListStudentByIDController(useCase *application.ListStudentById) *ListStudentByIDController {
	return &ListStudentByIDController{useCase: useCase}
}

func (lsid_c *ListStudentByIDController) Execute(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID no válido"})
		return
	}

	student, err := lsid_c.useCase.Execute(idInt)
	if err != nil {
		c.JSON(404, gin.H{"error": "Estudiante no encontrado"})
		return
	}

	c.JSON(200, student)
}
