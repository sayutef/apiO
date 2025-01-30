package controllers

import (
	"ApiRestAct1/src/students/application"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DeleteStudentController struct {
	useCase *application.DeleteStudent
}

func NewDeleteStudentController(useCase *application.DeleteStudent) *DeleteStudentController {
	return &DeleteStudentController{useCase: useCase}
}

func (ds_c *DeleteStudentController) Execute(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID no válido"})
		return
	}
	if err := ds_c.useCase.Execute(idInt); err != nil {
		c.JSON(500, gin.H{"error": "No se pudo eliminar el estudiante"})
		return
	}
	c.JSON(200, gin.H{"message": "Estudiante eliminado"})
}
