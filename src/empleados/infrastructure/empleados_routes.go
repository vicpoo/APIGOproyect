package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOproyect/src/empleados/application"
)

func RegisterEmpleadoRoutes(router *gin.Engine, createUC *application.CreateEmpleado, updateUC *application.UpdateEmpleado, deleteUC *application.DeleteEmpleado, viewUC *application.ViewEmpleado) {
	empleadoGroup := router.Group("/empleados")
	{
		empleadoGroup.POST("/", NewCreateEmpleadoController(createUC).Execute)
		empleadoGroup.PUT("/:id", NewUpdateEmpleadoController(updateUC).Execute)
		empleadoGroup.DELETE("/:id", NewDeleteEmpleadoController(deleteUC).Execute)
		empleadoGroup.GET("/:id", NewViewEmpleadoController(viewUC).Execute)
	}
}
