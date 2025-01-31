package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOproyect/src/empleados/application"
)

// ViewAllEmpleadosController es el controlador para obtener todos los empleados.
type ViewAllEmpleadosController struct {
	useCase *application.ViewAllEmpleados
}

// NewViewAllEmpleadosController crea una nueva instancia de ViewAllEmpleadosController.
func NewViewAllEmpleadosController(useCase *application.ViewAllEmpleados) *ViewAllEmpleadosController {
	return &ViewAllEmpleadosController{useCase: useCase}
}

// Execute maneja la solicitud HTTP para obtener todos los empleados.
func (vec *ViewAllEmpleadosController) Execute(c *gin.Context) {
	// Ejecutar el caso de uso
	empleados, err := vec.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los empleados"})
		return
	}

	// Devolver la lista de empleados en formato JSON
	c.JSON(http.StatusOK, empleados)
}
