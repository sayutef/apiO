package controllers

import (
	"ApiRestAct1/src/asignatures/application"
	"github.com/gin-gonic/gin"
)

type ListAsignatureController struct {
	useCase *application.ListAsignature
}

func NewListAsignatureController(useCase *application.ListAsignature) *ListAsignatureController {
	return &ListAsignatureController{useCase: useCase}
}

func (la_c *ListAsignatureController) Execute(c *gin.Context) {
	asignature, err := la_c.useCase.Execute()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener las asignaturas"})
		return
	}
	c.JSON(200, asignature)
}
