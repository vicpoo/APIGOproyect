package infrastructure

import (
	"github.com/vicpoo/APIGOproyect/src/equipos/application"
)

func InitEquipoDependencies() (
	*CreateEquipoController,
	*ViewEquipoController,
	*UpdateEquipoController,
	*DeleteEquipoController,
	*ViewAllEquiposController,
) {
	// Inicializar el repositorio
	repo := NewMysqlEquipoRepository()

	// Crear casos de uso
	createUseCase := application.NewCreateEquipoUseCase(repo)
	viewUseCase := application.NewViewEquipo(repo)
	updateUseCase := application.NewUpdateEquipo(repo)
	deleteUseCase := application.NewDeleteEquipoUseCase(repo)
	viewAllUseCase := application.NewViewAllEquipos(repo)

	// Crear controladores
	createController := NewCreateEquipoController(createUseCase)
	viewController := NewViewEquipoController(viewUseCase)
	updateController := NewUpdateEquipoController(updateUseCase)
	deleteController := NewDeleteEquipoController(deleteUseCase)
	viewAllController := NewViewAllEquiposController(viewAllUseCase)

	return createController, viewController, updateController, deleteController, viewAllController
}
