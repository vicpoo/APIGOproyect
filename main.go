package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOproyect/src/core"
	empleadosInfra "github.com/vicpoo/APIGOproyect/src/empleados/infrastructure" // Alias para empleados/infrastructure
	equiposInfra "github.com/vicpoo/APIGOproyect/src/equipos/infrastructure"     // Alias para equipos/infrastructure
)

func main() {
	// Iniciar la conexión a la base de datos
	core.InitDB()

	// Configurar el servidor HTTP
	r := gin.Default()

	// Inicializar el enrutador de empleados y configurar las rutas
	empleadosRouter := empleadosInfra.NewRouter(r) // Usar el alias empleadosInfra
	empleadosRouter.Run()

	// Inicializar el enrutador de equipos y configurar las rutas
	equiposRouter := equiposInfra.NewRouter(r) // Usar el alias equiposInfra
	equiposRouter.Run()

	// Iniciar el servidor
	err := r.Run(":8000")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
