package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOproyect/src/empleados/application"
)

type ViewAllEmpleadosController struct {
	useCase *application.ViewAllEmpleados
}

func NewViewAllEmpleadosController(useCase *application.ViewAllEmpleados) *ViewAllEmpleadosController {
	return &ViewAllEmpleadosController{useCase: useCase}
}

func (vec *ViewAllEmpleadosController) Execute(c *gin.Context) {
	empleados, err := vec.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los empleados"})
		return
	}

	c.JSON(http.StatusOK, empleados)
}
