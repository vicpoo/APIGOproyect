package main

import (
	"fmt"

	"github.com/vicpoo/APIGOproyect/src/core"
)

func main() {
	core.InitDB()
	fmt.Println("Servidor iniciado en el puerto 8000")
}
