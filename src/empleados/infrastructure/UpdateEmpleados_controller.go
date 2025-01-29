package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOproyect/src/empleados/application"
	"github.com/vicpoo/APIGOproyect/src/empleados/domain/entities"
)

type UpdateEmpleadoController struct {
	useCase *application.UpdateEmpleado
}

func NewUpdateEmpleadoController(useCase *application.UpdateEmpleado) *UpdateEmpleadoController {
	return &UpdateEmpleadoController{useCase: useCase}
}

func (uec *UpdateEmpleadoController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var empleado entities.Empleado
	if err := c.ShouldBindJSON(&empleado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uec.useCase.Execute(id, empleado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el empleado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Empleado actualizado exitosamente"})
}
