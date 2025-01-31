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
	createController, viewController, updateController, deleteController, viewAllController := InitEquipoDependencies()

	// Grupo de rutas para equipos
	equipoGroup := router.engine.Group("/equipos")
	{
		equipoGroup.POST("/", createController.Run)       // Crear equipo
		equipoGroup.GET("/:id", viewController.Execute)   // Ver equipo por ID
		equipoGroup.PUT("/:id", updateController.Execute) // Actualizar equipo
		equipoGroup.DELETE("/:id", deleteController.Run)  // Eliminar equipo
		equipoGroup.GET("/", viewAllController.Execute)   // Ver todos los equipos
	}
}
