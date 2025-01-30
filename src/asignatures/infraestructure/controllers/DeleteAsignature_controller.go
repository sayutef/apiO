package controllers

import (
	"ApiRestAct1/src/asignatures/application"
	"github.com/gin-gonic/gin"
	"strconv"
)

type DeleteAsignatureController struct {
	useCase *application.DeleteAsignature
}

func NewDeleteAsignatureController(useCase *application.DeleteAsignature) *DeleteAsignatureController {
	return &DeleteAsignatureController{useCase: useCase}
}

func (da_c *DeleteAsignatureController) Execute(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID no válido"})
		return
	}
	if err := da_c.useCase.Execute(idInt); err != nil {
		c.JSON(500, gin.H{"error": "No se pudo eliminar la asignatura"})
		return
	}
	c.JSON(200, gin.H{"message": "Asignatura eliminada"})
}
