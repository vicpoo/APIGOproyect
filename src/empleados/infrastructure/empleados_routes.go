package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func NewRouter(engine *gin.Engine) *Router {
	return &Router{
		engine: engine,
	}
}

func (router *Router) Run() {
	// Inicializar dependencias
	createController, viewController, updateController, deleteController, viewAllController := InitEmpleadoDependencies()

	// Grupo de rutas para empleados
	empleadoGroup := router.engine.Group("/empleados")
	{
		empleadoGroup.POST("/", createController.Run)       // Crear empleado
		empleadoGroup.GET("/:id", viewController.Execute)   // Ver empleado por ID
		empleadoGroup.PUT("/:id", updateController.Execute) // Actualizar empleado
		empleadoGroup.DELETE("/:id", deleteController.Run)  // Eliminar empleado
		empleadoGroup.GET("/", viewAllController.Execute)   // Ver todos los empleados
	}
}
