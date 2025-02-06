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
	createController, viewController, updateController, deleteController, viewAllController, GetEmpleadosByStatusController := InitEmpleadoDependencies()

	// Grupo de rutas para empleados
	empleadoGroup := router.engine.Group("/empleados")
	{
		empleadoGroup.POST("/", createController.Run)
		empleadoGroup.GET("/:id", viewController.Execute)
		empleadoGroup.PUT("/:id", updateController.Execute)
		empleadoGroup.DELETE("/:id", deleteController.Run)
		empleadoGroup.GET("/", viewAllController.Execute)
		empleadoGroup.GET("/status/:status", GetEmpleadosByStatusController.Run)
	}
}
