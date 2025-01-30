package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOproyect/src/core"
	"github.com/vicpoo/APIGOproyect/src/empleados/infrastructure"
)

func main() {
	// Iniciar la conexión a la base de datos
	core.InitDB()

	// Configurar el servidor HTTP
	r := gin.Default()

	// Inicializar el enrutador y configurar las rutas
	router := infrastructure.NewRouter(r)
	router.Run()

	// Iniciar el servidor
	err := r.Run(":8000")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
