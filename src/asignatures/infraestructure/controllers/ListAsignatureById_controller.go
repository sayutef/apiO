package controllers

import (
	"ApiRestAct1/src/asignatures/application"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ListAsignatureByIdController struct {
	useCase *application.ListAsignatureById
}

func NewListAsignatureByIdController(useCase *application.ListAsignatureById) *ListAsignatureByIdController {
	return &ListAsignatureByIdController{useCase: useCase}
}

func (laid_c *ListAsignatureByIdController) Execute(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "ID no válido"})
		return
	}

	asignature, err := laid_c.useCase.Execute(idInt)
	if err != nil {
		c.JSON(404, gin.H{"error": "Asignatura no encontrada"})
		return
	}

	c.JSON(200, asignature)
}
