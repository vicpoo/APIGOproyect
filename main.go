package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOproyect/core"
	"github.com/vicpoo/APIGOproyect/src/empleados/application"
	"github.com/vicpoo/APIGOproyect/src/empleados/domain" // Asegúrate de que este import esté presente
	"github.com/vicpoo/APIGOproyect/src/empleados/infrastructure"
)

func main() {
	// Iniciar la conexión a la base de datos
	core.InitDB()

	// Crear el repositorio de empleados
	empleadoRepo := domain.NewEmpleadoRepository(core.GetDB()) // Usar correctamente el paquete domain

	// Crear el caso de uso para crear empleados
	createEmpleado := application.NewCreateEmpleado(empleadoRepo)

	// Crear el controlador
	createEmpleadoController := infrastructure.NewCreateEmpleadoController(createEmpleado)

	// Configurar el servidor HTTP
	r := gin.Default()
	r.POST("/empleados", createEmpleadoController.Execute)

	// Iniciar el servidor
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
