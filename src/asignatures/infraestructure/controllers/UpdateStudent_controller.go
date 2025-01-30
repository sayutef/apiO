package controllers

import (
	"ApiRestAct1/src/asignatures/application"
	"ApiRestAct1/src/asignatures/domain/entities"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UpdateAsignatureController struct {
	useCase *application.UpdateAsignature
}

func NewUpdateAsignatureController(useCase *application.UpdateAsignature) *UpdateAsignatureController {
	return &UpdateAsignatureController{useCase: useCase}
}

func (us_c *UpdateAsignatureController) Execute(c *gin.Context) {
	var asignature entities.Asignature
	if err := c.ShouldBindJSON(&asignature); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID no válido"})
		return
	}
	asignature.ID = idInt
	if err := us_c.useCase.Execute(asignature); err != nil {
		c.JSON(500, gin.H{"error": "No se pudo actualizar el estudiante"})
		return
	}
	c.JSON(200, gin.H{"message": "Estudiante actualizado"})
}
