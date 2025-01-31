package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOproyect/src/equipos/application"
)

type ViewAllEquiposController struct {
	useCase *application.ViewAllEquipos
}

func NewViewAllEquiposController(useCase *application.ViewAllEquipos) *ViewAllEquiposController {
	return &ViewAllEquiposController{useCase: useCase}
}

func (vec *ViewAllEquiposController) Execute(c *gin.Context) {
	equipos, err := vec.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los equipos"})
		return
	}

	c.JSON(http.StatusOK, equipos)
}
