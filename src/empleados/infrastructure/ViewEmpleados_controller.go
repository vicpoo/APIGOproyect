package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOproyect/src/empleados/application"
)

type ViewEmpleadoController struct {
	useCase *application.ViewEmpleado
}

func NewViewEmpleadoController(useCase *application.ViewEmpleado) *ViewEmpleadoController {
	return &ViewEmpleadoController{useCase: useCase}
}

func (vec *ViewEmpleadoController) Execute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	empleado, err := vec.useCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Empleado no encontrado"})
		return
	}

	c.JSON(http.StatusOK, empleado)
}
