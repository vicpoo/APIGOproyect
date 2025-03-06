package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOproyect/src/core"
	empleadosInfra "github.com/vicpoo/APIGOproyect/src/empleados/infrastructure"
	equiposInfra "github.com/vicpoo/APIGOproyect/src/equipos/infrastructure"
)

func main() {
	core.InitDB()

	r := gin.Default()

	// Configuración de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, // Cambia esto según la URL de tu frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Configuración de rutas para empleados
	empleadosRouter := empleadosInfra.NewRouter(r)
	empleadosRouter.Run()

	// Configuración de rutas para equipos
	equiposRouter := equiposInfra.NewRouter(r)
	equiposRouter.Run()

	// Iniciar el servidor en el puerto 8000
	err := r.Run(":8000")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
