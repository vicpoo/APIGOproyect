package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOproyect/src/equipos/application"
	"github.com/vicpoo/APIGOproyect/src/equipos/domain/entities"
)

type CreateEquipoController struct {
	CreateEquipoUseCase *application.CreateEquipoUseCase
}

func NewCreateEquipoController(createEquipoUseCase *application.CreateEquipoUseCase) *CreateEquipoController {
	return &CreateEquipoController{
		CreateEquipoUseCase: createEquipoUseCase,
	}
}

func (ctrl *CreateEquipoController) Run(c *gin.Context) {
	var equipo entities.Equipo

	if errJSON := c.ShouldBindJSON(&equipo); errJSON != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Datos del equipo inválidos",
			"error":   errJSON.Error(),
		})
		return
	}

	equipoCreado, errAdd := ctrl.CreateEquipoUseCase.Run(&equipo)

	if errAdd != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al agregar el equipo",
			"error":   errAdd.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "El equipo ha sido agregado",
		"equipo":  equipoCreado,
	})
}
