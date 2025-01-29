package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOproyect/src/empleados/application"
)

type DeleteEmpleadoController struct {
	useCase *application.DeleteEmpleado
}

func NewDeleteEmpleadoController(useCase *application.DeleteEmpleado) *DeleteEmpleadoController {
	return &DeleteEmpleadoController{useCase: useCase}
}

func (dec *DeleteEmpleadoController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = dec.useCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el empleado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Empleado eliminado exitosamente"})
}
