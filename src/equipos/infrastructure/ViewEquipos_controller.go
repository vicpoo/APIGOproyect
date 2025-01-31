package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOproyect/src/equipos/application"
)

type ViewEquipoController struct {
	useCase *application.ViewEquipo
}

func NewViewEquipoController(useCase *application.ViewEquipo) *ViewEquipoController {
	return &ViewEquipoController{useCase: useCase}
}

func (vec *ViewEquipoController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	equipo, err := vec.useCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Equipo no encontrado"})
		return
	}

	c.JSON(http.StatusOK, equipo)
}
