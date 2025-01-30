package infrastructure

import (
	"github.com/vicpoo/APIGOproyect/src/empleados/application"
)

func InitEmpleadoDependencies() (
	*CreateEmpleadoController,
	*ViewEmpleadoController,
	*UpdateEmpleadoController,
	*DeleteEmpleadoController,
) {
	// Inicializar el repositorio
	repo := NewMysqlEmpleadoRepository()

	// Crear casos de uso
	createUseCase := application.NewCreateEmpleadoUseCase(repo)
	viewUseCase := application.NewViewEmpleado(repo)
	updateUseCase := application.NewUpdateEmpleado(repo)
	deleteUseCase := application.NewDeleteEmpleadoUseCase(repo)

	// Crear controladores
	createController := NewCreateEmpleadoController(createUseCase)
	viewController := NewViewEmpleadoController(viewUseCase)
	updateController := NewUpdateEmpleadoController(updateUseCase)
	deleteController := NewDeleteEmpleadoController(deleteUseCase)

	return createController, viewController, updateController, deleteController
}
