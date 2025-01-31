package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOproyect/src/equipos/application"
	"github.com/vicpoo/APIGOproyect/src/equipos/domain/entities"
)

type UpdateEquipoController struct {
	useCase *application.UpdateEquipo
}

func NewUpdateEquipoController(useCase *application.UpdateEquipo) *UpdateEquipoController {
	return &UpdateEquipoController{useCase: useCase}
}

func (uec *UpdateEquipoController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var equipo entities.Equipo
	if err := c.ShouldBindJSON(&equipo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uec.useCase.Execute(id, equipo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el equipo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Equipo actualizado exitosamente"})
}
