package controllers

import (
	"ApiRestAct1/src/asignatures/application"
	"ApiRestAct1/src/asignatures/domain/entities"
	"github.com/gin-gonic/gin"
)

type CreateAsignatureController struct {
	useCase *application.CreateAsignature
}

func NewCreateAsignatureController(useCase *application.CreateAsignature) *CreateAsignatureController {
	return &CreateAsignatureController{useCase: useCase}
}

func (cs_a *CreateAsignatureController) Execute(c *gin.Context) {
	var asignature entities.Asignature
	if err := c.ShouldBindJSON(&asignature); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}
	if err := cs_a.useCase.Execute(asignature); err != nil {
		c.JSON(500, gin.H{"error": "No se pudo crear el asignatura"})
		return
	}
	c.JSON(201, gin.H{"message": "Asignatura creado correctamente"})
}
