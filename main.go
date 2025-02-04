package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/APIGOproyect/src/core"
	empleadosInfra "github.com/vicpoo/APIGOproyect/src/empleados/infrastructure"
	equiposInfra "github.com/vicpoo/APIGOproyect/src/equipos/infrastructure"
)

func main() {

	core.InitDB()

	r := gin.Default()

	empleadosRouter := empleadosInfra.NewRouter(r)
	empleadosRouter.Run()

	equiposRouter := equiposInfra.NewRouter(r)
	equiposRouter.Run()

	err := r.Run(":8000")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
