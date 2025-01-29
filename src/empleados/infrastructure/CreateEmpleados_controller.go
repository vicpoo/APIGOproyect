package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOproyect/src/empleados/application"
	"github.com/vicpoo/APIGOproyect/src/empleados/domain"
)

type CreateEmpleadoController struct {
	useCase *application.CreateEmpleado
}

func NewCreateEmpleadoController(useCase *application.CreateEmpleado) *CreateEmpleadoController {
	return &CreateEmpleadoController{useCase: useCase}
}

func (cec *CreateEmpleadoController) Execute(c *gin.Context) {
	var empleado domain.Empleado
	if err := c.ShouldBindJSON(&empleado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := cec.useCase.Execute(empleado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear el empleado"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Empleado creado exitosamente"})
}
