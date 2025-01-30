package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOproyect/src/empleados/application"
	"github.com/vicpoo/APIGOproyect/src/empleados/domain/entities"
)

type CreateEmpleadoController struct {
	CreateEmpleadoUseCase *application.CreateEmpleadoUseCase
}

func NewCreateEmpleadoController(createEmpleadoUseCase *application.CreateEmpleadoUseCase) *CreateEmpleadoController {
	return &CreateEmpleadoController{
		CreateEmpleadoUseCase: createEmpleadoUseCase,
	}
}

func (ctrl *CreateEmpleadoController) Run(c *gin.Context) {
	var empleado entities.Empleado

	if errJSON := c.ShouldBindJSON(&empleado); errJSON != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos del empleado inválidos",
			"error":   errJSON.Error(),
		})
		return
	}

	empleadoCreado, errAdd := ctrl.CreateEmpleadoUseCase.Run(&empleado)

	if errAdd != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al agregar el empleado",
			"error":   errAdd.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "El empleado ha sido agregado",
		"empleado": empleadoCreado,
	})
}
