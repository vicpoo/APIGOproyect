package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	application "github.com/vicpoo/APIGOproyect/src/empleados/application"
)

type GetEmpleadosByStatusController struct {
	getByStatusUseCase *application.GetEmpleadosByStatusUseCase
}

func NewGetEmpleadosByStatusController(getByStatusUseCase *application.GetEmpleadosByStatusUseCase) *GetEmpleadosByStatusController {
	return &GetEmpleadosByStatusController{
		getByStatusUseCase: getByStatusUseCase,
	}
}

func (ctrl *GetEmpleadosByStatusController) Run(c *gin.Context) {
	statusParam := c.Param("status")
	status, err := strconv.ParseBool(statusParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Estado inválido",
			"error":   err.Error(),
		})
		return
	}

	empleados, err := ctrl.getByStatusUseCase.Run(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error al obtener empleados por estado",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"empleados": empleados,
	})
}
