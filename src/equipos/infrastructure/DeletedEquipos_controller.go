package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	application "github.com/vicpoo/APIGOproyect/src/equipos/application"
)

type DeleteEquipoController struct {
	deleteUseCase *application.DeleteEquipoUseCase
}

func NewDeleteEquipoController(deleteUseCase *application.DeleteEquipoUseCase) *DeleteEquipoController {
	return &DeleteEquipoController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteEquipoController) Run(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	errDelete := ctrl.deleteUseCase.Run(id)
	if errDelete != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "No se pudo eliminar el equipo",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Equipo eliminado exitosamente",
	})
}
