package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	application "github.com/vicpoo/APIGOproyect/src/empleados/application"
)

type DeleteEmpleadoController struct {
	deleteUseCase *application.DeleteEmpleadoUseCase
}

func NewDeleteEmpleadoController(deleteUseCase *application.DeleteEmpleadoUseCase) *DeleteEmpleadoController {
	return &DeleteEmpleadoController{
		deleteUseCase: deleteUseCase,
	}
}

func (ctrl *DeleteEmpleadoController) Run(c *gin.Context) {
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
			"message": "No se pudo eliminar el empleado",
			"error":   errDelete.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Empleado eliminado exitosamente",
	})
}
